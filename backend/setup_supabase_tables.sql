-- Quick Supabase Setup Script
-- Run this in your Supabase SQL Editor

-- 1. Create users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    phone VARCHAR(20),
    coin_balance INTEGER DEFAULT 50,
    role VARCHAR(20) DEFAULT 'user',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 2. Create series table
CREATE TABLE IF NOT EXISTS series (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    author VARCHAR(255) NOT NULL,
    category VARCHAR(100),
    is_premium BOOLEAN DEFAULT false,
    total_episodes INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 3. Create episodes table
CREATE TABLE IF NOT EXISTS episodes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    series_id UUID REFERENCES series(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    audio_url TEXT,
    duration INTEGER DEFAULT 0,
    episode_number INTEGER NOT NULL,
    coin_price INTEGER DEFAULT 10,
    is_locked BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 4. Create coin_bundles table
CREATE TABLE IF NOT EXISTS coin_bundles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    coins INTEGER NOT NULL,
    price INTEGER NOT NULL,
    currency VARCHAR(3) DEFAULT 'INR',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 5. Insert sample data
INSERT INTO series (title, description, author, category, is_premium) VALUES
('The Mystery of the Lost City', 'An adventure series about discovering ancient civilizations', 'John Smith', 'Adventure', true),
('Tech Talk Daily', 'Daily insights into the world of technology', 'Sarah Johnson', 'Technology', false)
ON CONFLICT DO NOTHING;

-- 6. Insert sample episodes
INSERT INTO episodes (series_id, title, description, audio_url, episode_number, coin_price, is_locked) 
SELECT 
    s.id,
    'Episode 1: The Discovery',
    'Our journey begins with an ancient map',
    'https://example.com/audio/ep1.mp3',
    1,
    10,
    true
FROM series s WHERE s.title = 'The Mystery of the Lost City'
ON CONFLICT DO NOTHING;

INSERT INTO episodes (series_id, title, description, audio_url, episode_number, coin_price, is_locked) 
SELECT 
    s.id,
    'Episode 2: The Temple',
    'We enter the mysterious temple',
    'https://example.com/audio/ep2.mp3',
    2,
    15,
    true
FROM series s WHERE s.title = 'The Mystery of the Lost City'
ON CONFLICT DO NOTHING;

-- 7. Insert sample coin bundles
INSERT INTO coin_bundles (name, coins, price, currency) VALUES
('Starter Pack', 50, 500, 'INR'),
('Popular Pack', 100, 900, 'INR'),
('Premium Pack', 200, 1600, 'INR')
ON CONFLICT DO NOTHING;

-- 8. Test queries to verify setup
-- Test 1: Count all series
SELECT 'Series Count' as test, COUNT(*) as result FROM series;

-- Test 2: Count all episodes
SELECT 'Episodes Count' as test, COUNT(*) as result FROM episodes;

-- Test 3: Count all coin bundles
SELECT 'Coin Bundles Count' as test, COUNT(*) as result FROM coin_bundles;

-- Test 4: Get series with episode count
SELECT 
    s.title,
    s.author,
    s.category,
    COUNT(e.id) as episode_count
FROM series s
LEFT JOIN episodes e ON s.id = e.series_id
GROUP BY s.id, s.title, s.author, s.category;

-- Test 5: Get episodes with series info
SELECT 
    e.title as episode_title,
    s.title as series_title,
    e.episode_number,
    e.coin_price,
    e.is_locked
FROM episodes e
JOIN series s ON e.series_id = s.id
ORDER BY s.title, e.episode_number; 