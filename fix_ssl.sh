#!/bin/bash

echo "🔧 Adding SSL Parameters to Supabase Connection"
echo "==============================================="
echo ""

# Create a backup of the current .env file
cp .env .env.backup.ssl
echo "✅ Created backup: .env.backup.ssl"

# Update the SUPABASE_URL with SSL parameters
sed -i '' 's|SUPABASE_URL=postgresql://postgres:hari%409944110859@db.mhbcihpkcetbzdrzciqe.supabase.co:5432/postgres|SUPABASE_URL=postgresql://postgres:hari%409944110859@db.mhbcihpkcetbzdrzciqe.supabase.co:5432/postgres?sslmode=require|' .env

echo "✅ Updated SUPABASE_URL with SSL parameters"
echo ""
echo "🔗 Testing connection..."
echo ""

go run test_connection.go 