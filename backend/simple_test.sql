-- Simple Test Table
-- Run this in your Supabase SQL Editor

-- Create a simple names table
CREATE TABLE IF NOT EXISTS names (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    age INTEGER,
    city VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Insert some sample data
INSERT INTO names (name, age, city) VALUES
('John Doe', 25, 'New York'),
('Jane Smith', 30, 'Los Angeles'),
('Bob Johnson', 35, 'Chicago')
ON CONFLICT DO NOTHING;

-- Test query to see the data
SELECT * FROM names ORDER BY created_at; 