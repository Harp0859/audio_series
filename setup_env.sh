#!/bin/bash

echo "🎧 Audio Series App - Environment Setup"
echo "======================================="
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
echo "🔧 Environment Configuration Options:"
echo "===================================="
echo ""
echo "You have two options for Supabase configuration:"
echo ""
echo "Option 1: Direct Database Connection (Recommended)"
echo "  SUPABASE_URL=postgresql://postgres:[PASSWORD]@db.[PROJECT-REF].supabase.co:5432/postgres"
echo ""
echo "Option 2: REST API URL + Database Password"
echo "  SUPABASE_URL=https://your-project.supabase.co"
echo "  SUPABASE_DB_PASSWORD=your_database_password"
echo ""
echo "To configure your environment:"
echo "1. Edit backend/.env with your actual Supabase credentials"
echo "2. Make sure to set SUPABASE_DB_PASSWORD if using REST API URL format"
echo "3. Run: cd backend && go run test_connection.go"
echo ""
echo "📚 For detailed instructions, check docs/SUPABASE_SETUP.md"
echo ""
echo "🎉 Setup complete! Happy coding!" 