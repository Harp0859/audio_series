-- Audio Series App Database Schema
-- This file contains the complete database schema for the audio series application

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    avatar_url TEXT,
    coin_balance INTEGER DEFAULT 0 NOT NULL,
    role VARCHAR(20) DEFAULT 'user' CHECK (role IN ('user', 'admin')),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Series table
CREATE TABLE series (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    cover_image TEXT,
    author VARCHAR(255) NOT NULL,
    category VARCHAR(100),
    is_premium BOOLEAN DEFAULT false,
    total_episodes INTEGER DEFAULT 0,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Episodes table
CREATE TABLE episodes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    series_id UUID REFERENCES series(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    audio_url TEXT NOT NULL,
    duration INTEGER DEFAULT 0, -- in seconds
    episode_number INTEGER NOT NULL,
    coin_price INTEGER DEFAULT 0,
    is_locked BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(series_id, episode_number)
);

-- Purchases table
CREATE TABLE purchases (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    episode_id UUID REFERENCES episodes(id) ON DELETE CASCADE,
    series_id UUID REFERENCES series(id) ON DELETE CASCADE,
    type VARCHAR(20) NOT NULL CHECK (type IN ('episode', 'series', 'coins')),
    amount INTEGER NOT NULL, -- coins spent
    payment_id VARCHAR(255),
    status VARCHAR(20) DEFAULT 'completed' CHECK (status IN ('completed', 'pending', 'failed')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CHECK (
        (episode_id IS NOT NULL AND series_id IS NULL) OR
        (episode_id IS NULL AND series_id IS NOT NULL) OR
        (episode_id IS NULL AND series_id IS NULL)
    )
);

-- Coin transactions table
CREATE TABLE coin_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(20) NOT NULL CHECK (type IN ('purchase', 'welcome', 'refund', 'admin')),
    amount INTEGER NOT NULL, -- positive for credit, negative for debit
    balance INTEGER NOT NULL, -- balance after transaction
    description TEXT,
    reference_id VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Payments table
CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    amount INTEGER NOT NULL, -- in smallest currency unit
    currency VARCHAR(3) NOT NULL,
    coins INTEGER NOT NULL,
    gateway VARCHAR(20) NOT NULL CHECK (gateway IN ('razorpay', 'paystack')),
    gateway_ref VARCHAR(255) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'failed')),
    payment_data JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Coin bundles table
CREATE TABLE coin_bundles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    coins INTEGER NOT NULL,
    price INTEGER NOT NULL, -- in smallest currency unit
    currency VARCHAR(3) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Indexes for better performance
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_series_created_by ON series(created_by);
CREATE INDEX idx_episodes_series_id ON episodes(series_id);
CREATE INDEX idx_episodes_episode_number ON episodes(series_id, episode_number);
CREATE INDEX idx_purchases_user_id ON purchases(user_id);
CREATE INDEX idx_purchases_episode_id ON purchases(episode_id);
CREATE INDEX idx_purchases_series_id ON purchases(series_id);
CREATE INDEX idx_coin_transactions_user_id ON coin_transactions(user_id);
CREATE INDEX idx_payments_user_id ON payments(user_id);
CREATE INDEX idx_payments_gateway_ref ON payments(gateway_ref);

-- Triggers to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_series_updated_at BEFORE UPDATE ON series
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_episodes_updated_at BEFORE UPDATE ON episodes
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_payments_updated_at BEFORE UPDATE ON payments
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Function to update series total_episodes count
CREATE OR REPLACE FUNCTION update_series_episode_count()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        UPDATE series SET total_episodes = total_episodes + 1 WHERE id = NEW.series_id;
        RETURN NEW;
    ELSIF TG_OP = 'DELETE' THEN
        UPDATE series SET total_episodes = total_episodes - 1 WHERE id = OLD.series_id;
        RETURN OLD;
    END IF;
    RETURN NULL;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_series_episode_count_trigger
    AFTER INSERT OR DELETE ON episodes
    FOR EACH ROW EXECUTE FUNCTION update_series_episode_count();

-- Insert default coin bundles
INSERT INTO coin_bundles (name, coins, price, currency, is_active) VALUES
('50 Coins', 50, 5000, 'INR', true),
('120 Coins', 120, 9900, 'INR', true),
('250 Coins', 250, 19900, 'INR', true),
('500 Coins', 500, 39900, 'INR', true),
('50 Coins', 50, 5000, 'NGN', true),
('120 Coins', 120, 9900, 'NGN', true),
('250 Coins', 250, 19900, 'NGN', true),
('500 Coins', 500, 39900, 'NGN', true);

-- Insert sample data for testing
INSERT INTO users (email, first_name, last_name, coin_balance, role) VALUES
('admin@audioseries.com', 'Admin', 'User', 1000, 'admin'),
('user@example.com', 'John', 'Doe', 100, 'user');

INSERT INTO series (title, description, cover_image, author, category, is_premium, total_episodes, created_by) VALUES
('Forbidden Nights', 'A thrilling audio series about mystery and suspense', 'https://example.com/cover1.jpg', 'Jane Smith', 'Mystery', true, 10, (SELECT id FROM users WHERE email = 'admin@audioseries.com')),
('Urban Legends', 'Modern urban legends brought to life', 'https://example.com/cover2.jpg', 'Mike Johnson', 'Horror', false, 8, (SELECT id FROM users WHERE email = 'admin@audioseries.com'));

INSERT INTO episodes (series_id, title, description, audio_url, duration, episode_number, coin_price, is_locked) VALUES
((SELECT id FROM series WHERE title = 'Forbidden Nights'), 'Episode 1: The Beginning', 'The story begins with a mysterious discovery', 'https://example.com/audio1.mp3', 1800, 1, 10, true),
((SELECT id FROM series WHERE title = 'Forbidden Nights'), 'Episode 2: The Investigation', 'The plot thickens as clues are uncovered', 'https://example.com/audio2.mp3', 1800, 2, 15, true),
((SELECT id FROM series WHERE title = 'Urban Legends'), 'Episode 1: The Legend Begins', 'The first urban legend comes to life', 'https://example.com/audio3.mp3', 1200, 1, 5, false); 