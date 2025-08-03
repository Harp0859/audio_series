# Supabase Database Schema Setup

## Step 1: Access Your Supabase Project

1. Go to your Supabase project dashboard: https://supabase.com/dashboard
2. Select your project: `mhbcihpkcetbzdrzciqe`

## Step 2: Set Up Database Schema

1. In your Supabase dashboard, go to **SQL Editor**
2. Copy and paste the entire contents of `database/schema.sql` into the SQL editor
3. Click **Run** to execute the schema

## Step 3: Configure Row Level Security (RLS)

After running the schema, you'll need to set up Row Level Security policies. Here are the basic policies you should add:

### Users Table Policies

```sql
-- Enable RLS on users table
ALTER TABLE users ENABLE ROW LEVEL SECURITY;

-- Users can read their own profile
CREATE POLICY "Users can view own profile" ON users
    FOR SELECT USING (auth.uid() = id);

-- Users can update their own profile
CREATE POLICY "Users can update own profile" ON users
    FOR UPDATE USING (auth.uid() = id);

-- Allow service role to manage all users
CREATE POLICY "Service role can manage all users" ON users
    FOR ALL USING (auth.role() = 'service_role');
```

### Series Table Policies

```sql
-- Enable RLS on series table
ALTER TABLE series ENABLE ROW LEVEL SECURITY;

-- Anyone can read series
CREATE POLICY "Anyone can read series" ON series
    FOR SELECT USING (true);

-- Only service role can manage series
CREATE POLICY "Service role can manage series" ON series
    FOR ALL USING (auth.role() = 'service_role');
```

### Episodes Table Policies

```sql
-- Enable RLS on episodes table
ALTER TABLE episodes ENABLE ROW LEVEL SECURITY;

-- Anyone can read episodes
CREATE POLICY "Anyone can read episodes" ON episodes
    FOR SELECT USING (true);

-- Only service role can manage episodes
CREATE POLICY "Service role can manage episodes" ON episodes
    FOR ALL USING (auth.role() = 'service_role');
```

### Purchases Table Policies

```sql
-- Enable RLS on purchases table
ALTER TABLE purchases ENABLE ROW LEVEL SECURITY;

-- Users can read their own purchases
CREATE POLICY "Users can view own purchases" ON purchases
    FOR SELECT USING (auth.uid() = user_id);

-- Users can create their own purchases
CREATE POLICY "Users can create own purchases" ON purchases
    FOR INSERT WITH CHECK (auth.uid() = user_id);

-- Service role can manage all purchases
CREATE POLICY "Service role can manage all purchases" ON purchases
    FOR ALL USING (auth.role() = 'service_role');
```

### Coin Transactions Table Policies

```sql
-- Enable RLS on coin_transactions table
ALTER TABLE coin_transactions ENABLE ROW LEVEL SECURITY;

-- Users can read their own transactions
CREATE POLICY "Users can view own transactions" ON coin_transactions
    FOR SELECT USING (auth.uid() = user_id);

-- Service role can manage all transactions
CREATE POLICY "Service role can manage all transactions" ON coin_transactions
    FOR ALL USING (auth.role() = 'service_role');
```

### Payments Table Policies

```sql
-- Enable RLS on payments table
ALTER TABLE payments ENABLE ROW LEVEL SECURITY;

-- Users can read their own payments
CREATE POLICY "Users can view own payments" ON payments
    FOR SELECT USING (auth.uid() = user_id);

-- Users can create their own payments
CREATE POLICY "Users can create own payments" ON payments
    FOR INSERT WITH CHECK (auth.uid() = user_id);

-- Service role can manage all payments
CREATE POLICY "Service role can manage all payments" ON payments
    FOR ALL USING (auth.role() = 'service_role');
```

### Coin Bundles Table Policies

```sql
-- Enable RLS on coin_bundles table
ALTER TABLE coin_bundles ENABLE ROW LEVEL SECURITY;

-- Anyone can read coin bundles
CREATE POLICY "Anyone can read coin bundles" ON coin_bundles
    FOR SELECT USING (true);

-- Only service role can manage coin bundles
CREATE POLICY "Service role can manage coin bundles" ON coin_bundles
    FOR ALL USING (auth.role() = 'service_role');
```

## Step 4: Test the Connection

1. Start the backend server:
   ```bash
   cd backend
   go run cmd/server/main.go
   ```

2. Open your browser and go to: `http://localhost:3003`

3. The frontend should now be able to connect to Supabase and handle authentication.

## Step 5: Add Sample Data (Optional)

You can add some sample data to test the application:

```sql
-- Insert sample series
INSERT INTO series (title, description, author, category, is_premium) VALUES
('The Mystery of the Lost City', 'An adventure series about discovering ancient civilizations', 'John Smith', 'Adventure', true),
('Tech Talk Daily', 'Daily insights into the world of technology', 'Sarah Johnson', 'Technology', false);

-- Insert sample episodes
INSERT INTO episodes (series_id, title, description, audio_url, episode_number, coin_price, is_locked) VALUES
((SELECT id FROM series WHERE title = 'The Mystery of the Lost City'), 'Episode 1: The Discovery', 'Our journey begins with an ancient map', 'https://example.com/audio/ep1.mp3', 1, 10, true),
((SELECT id FROM series WHERE title = 'The Mystery of the Lost City'), 'Episode 2: The Temple', 'We enter the mysterious temple', 'https://example.com/audio/ep2.mp3', 2, 15, true),
((SELECT id FROM series WHERE title = 'Tech Talk Daily'), 'AI Revolution', 'How AI is changing the world', 'https://example.com/audio/ai.mp3', 1, 0, false);

-- Insert sample coin bundles
INSERT INTO coin_bundles (name, coins, price, currency, is_active) VALUES
('Starter Pack', 50, 500, 'INR', true),
('Popular Pack', 100, 900, 'INR', true),
('Premium Pack', 200, 1600, 'INR', true);
```

## Troubleshooting

If you encounter issues:

1. **Authentication errors**: Make sure your Supabase keys are correct
2. **Database errors**: Check that the schema was created successfully
3. **RLS errors**: Ensure all RLS policies are in place
4. **CORS errors**: Check that your Supabase project allows requests from your domain

The application is now set up to use Supabase for authentication and data storage! 