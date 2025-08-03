#!/bin/bash

echo "🎧 Audio Series App - Supabase Setup"
echo "====================================="
echo ""

# Check if .env file exists
if [ ! -f "backend/.env" ]; then
    echo "📝 Creating .env file from template..."
    cp backend/env.example backend/.env
    echo "✅ Created backend/.env"
else
    echo "📝 .env file already exists"
fi

echo ""
echo "🔧 Supabase Configuration Instructions:"
echo "======================================"
echo ""
echo "1. Go to your Supabase project dashboard:"
echo "   https://supabase.com/dashboard"
echo ""
echo "2. Navigate to Settings > Database"
echo ""
echo "3. Find your database connection string. It should look like:"
echo "   postgresql://postgres:[YOUR-PASSWORD]@db.[YOUR-PROJECT-REF].supabase.co:5432/postgres"
echo ""
echo "4. Update backend/.env with your Supabase credentials:"
echo "   - SUPABASE_URL: Your database connection string"
echo "   - SUPABASE_ANON_KEY: Your project's anon key"
echo "   - SUPABASE_SERVICE_ROLE_KEY: Your project's service role key"
echo ""
echo "5. Run the database schema:"
echo "   - Copy the contents of database/schema.sql"
echo "   - Go to your Supabase SQL Editor"
echo "   - Paste and execute the schema"
echo ""
echo "6. Test the connection:"
echo "   cd backend && go run test_connection.go"
echo ""
echo "7. Start the server:"
echo "   cd backend && go run cmd/server/main.go"
echo ""
echo "📚 For more help, check the README.md file"
echo ""
echo "🎉 Setup complete! Happy coding!" 