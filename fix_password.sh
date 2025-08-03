#!/bin/bash

echo "🔧 Fixing Supabase Password Encoding"
echo "===================================="
echo ""

# Create a backup of the current .env file
cp .env .env.backup
echo "✅ Created backup: .env.backup"

# Update the SUPABASE_URL with properly encoded password
sed -i '' 's|SUPABASE_URL=postgresql://postgres:hari@9944110859@|SUPABASE_URL=postgresql://postgres:hari%409944110859@|' .env

echo "✅ Updated SUPABASE_URL with properly encoded password"
echo ""
echo "🔗 Testing connection..."
echo ""

go run test_connection.go 