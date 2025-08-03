#!/bin/bash

echo "🎧 Audio Series App - Supabase Configuration"
echo "============================================"
echo ""

# Check if .env file exists
if [ ! -f "backend/.env" ]; then
    echo "📝 Creating .env file from template..."
    cp backend/env.example backend/.env
    echo "✅ Created backend/.env"
fi

echo ""
echo "🔧 Current Configuration:"
echo "========================"
echo ""

# Show current SUPABASE_URL
CURRENT_URL=$(grep "^SUPABASE_URL=" backend/.env | cut -d'=' -f2)
echo "Current SUPABASE_URL: $CURRENT_URL"

# Show if SUPABASE_DB_PASSWORD is set
if grep -q "^SUPABASE_DB_PASSWORD=" backend/.env; then
    echo "SUPABASE_DB_PASSWORD: [SET]"
else
    echo "SUPABASE_DB_PASSWORD: [NOT SET]"
fi

echo ""
echo "📝 To configure your Supabase connection:"
echo "========================================="
echo ""
echo "1. Edit backend/.env and update these values:"
echo ""
echo "   For REST API URL format (your current format):"
echo "   SUPABASE_URL=https://your-project.supabase.co"
echo "   SUPABASE_ANON_KEY=your_anon_key"
echo "   SUPABASE_SERVICE_ROLE_KEY=your_service_role_key"
echo "   SUPABASE_DB_PASSWORD=your_database_password"
echo ""
echo "   OR for direct database connection:"
echo "   SUPABASE_URL=postgresql://postgres:[PASSWORD]@db.[PROJECT-REF].supabase.co:5432/postgres"
echo "   SUPABASE_ANON_KEY=your_anon_key"
echo "   SUPABASE_SERVICE_ROLE_KEY=your_service_role_key"
echo ""
echo "2. Test the connection:"
echo "   cd backend && go run test_connection.go"
echo ""
echo "3. Start the server:"
echo "   cd backend && go run cmd/server/main.go"
echo ""
echo "📚 Need help? Check docs/SUPABASE_SETUP.md" 