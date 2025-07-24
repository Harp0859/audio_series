#!/bin/bash

echo "🚀 Setting up Audio Series App..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21+ first."
    exit 1
fi

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js 18+ first."
    exit 1
fi

echo "📦 Installing backend dependencies..."
cd backend
go mod tidy
echo "✅ Backend dependencies installed"

echo "📦 Installing frontend dependencies..."
cd ../frontend
npm install
echo "✅ Frontend dependencies installed"

echo "📦 Installing mobile dependencies..."
cd ../mobile
npm install
echo "✅ Mobile dependencies installed"

echo "🔧 Setting up environment files..."
cd ..

# Create .env files if they don't exist
if [ ! -f "backend/.env" ]; then
    echo "📝 Creating backend .env file..."
    cp backend/env.example backend/.env
    echo "⚠️  Please update backend/.env with your Supabase credentials"
fi

if [ ! -f "frontend/.env" ]; then
    echo "📝 Creating frontend .env file..."
    cat > frontend/.env << EOF
REACT_APP_API_URL=http://localhost:8080
REACT_APP_SUPABASE_URL=your_supabase_url
REACT_APP_SUPABASE_ANON_KEY=your_supabase_anon_key
EOF
    echo "⚠️  Please update frontend/.env with your configuration"
fi

if [ ! -f "mobile/.env" ]; then
    echo "📝 Creating mobile .env file..."
    cat > mobile/.env << EOF
REACT_APP_API_URL=http://localhost:8080
REACT_APP_SUPABASE_URL=your_supabase_url
REACT_APP_SUPABASE_ANON_KEY=your_supabase_anon_key
EOF
    echo "⚠️  Please update mobile/.env with your configuration"
fi

echo "📊 Setting up database..."
echo "ℹ️  Please run the database schema in your Supabase project:"
echo "   - Go to your Supabase dashboard"
echo "   - Navigate to SQL Editor"
echo "   - Copy and paste the contents of database/schema.sql"
echo "   - Execute the script"

echo ""
echo "🎉 Setup complete!"
echo ""
echo "📋 Next steps:"
echo "1. Update environment files with your credentials"
echo "2. Set up your Supabase database using database/schema.sql"
echo "3. Start the backend: cd backend && go run cmd/server/main.go"
echo "4. Start the frontend: cd frontend && npm start"
echo "5. Start the mobile app: cd mobile && npx react-native run-ios"
echo ""
echo "🌐 Backend will run on: http://localhost:3003"
echo "🌐 Frontend will run on: http://localhost:3000"
echo ""
echo "📚 For more information, see README.md" 