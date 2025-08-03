# 🎧 Audio Series App

A complete, scalable audio series mobile and web application with coin-based unlocking system, payment integration, and admin dashboard.

## 🚀 Features

### User Features
- **User Authentication**: Signup/login with email, phone, and Google OAuth
- **Coin System**: Virtual currency for unlocking episodes
- **Audio Streaming**: Stream or download unlocked episodes
- **Series Browsing**: Browse available series and episodes
- **Payment Integration**: Buy coins via Razorpay (India) and Paystack (Nigeria)

### Admin Features
- **Series Management**: Create and manage audio series
- **Episode Management**: Add episodes with audio files and metadata
- **Pricing Control**: Set coin prices for episodes and bundles
- **Analytics Dashboard**: View user purchases and earnings

## 🛠 Tech Stack

- **Backend**: Go with Gin framework
- **Database**: Supabase PostgreSQL (Direct Connection)
- **Authentication**: JWT with bcrypt
- **Storage**: Supabase Storage for audio files
- **Frontend**: React.js (Web) + React Native (Mobile)
- **Payment**: Razorpay (India) + Paystack (Nigeria)
- **Language**: TypeScript (Frontend) + Go (Backend)

## 📁 Project Structure

```
audio-series-app/
├── backend/                 # Go backend API
│   ├── cmd/
│   ├── internal/
│   ├── pkg/
│   └── go.mod
├── frontend/               # React web app
│   ├── src/
│   ├── public/
│   └── package.json
├── mobile/                 # React Native mobile app
│   ├── src/
│   └── package.json
├── database/              # Database schemas and migrations
└── docs/                  # Documentation
```

## 🚀 Quick Start

### 1. Clone and Setup
```bash
git clone <repository-url>
cd audio-series-app
./setup_supabase.sh
```

### 2. Configure Supabase

#### Create a Supabase Project
1. Go to [Supabase Dashboard](https://supabase.com/dashboard)
2. Create a new project
3. Wait for the project to be ready

#### Get Database Connection String
1. In your Supabase project, go to **Settings > Database**
2. Find the **Connection string** section
3. Copy the **Direct connection** string that looks like:
   ```
   postgresql://postgres:[YOUR-PASSWORD]@db.[YOUR-PROJECT-REF].supabase.co:5432/postgres
   ```

#### Update Environment Variables
1. Edit `backend/.env`:
   ```env
   SUPABASE_URL=postgresql://postgres:[YOUR-PASSWORD]@db.[YOUR-PROJECT-REF].supabase.co:5432/postgres
   SUPABASE_ANON_KEY=your_supabase_anon_key
   SUPABASE_SERVICE_ROLE_KEY=your_supabase_service_role_key
   JWT_SECRET=your_secure_jwt_secret
   ```

#### Setup Database Schema
1. Go to your Supabase project's **SQL Editor**
2. Copy the contents of `database/schema.sql`
3. Paste and execute the SQL
4. Verify the tables are created in the **Table Editor**

### 3. Install Dependencies
```bash
# Backend
cd backend
go mod tidy

# Frontend
cd ../frontend
npm install

# Mobile
cd ../mobile
npm install
```

### 4. Start Applications
```bash
# Backend (port 3003)
cd backend
go run cmd/server/main.go

# Frontend (port 3004)
cd ../frontend
PORT=3004 npm start

# Mobile
cd ../mobile
npx react-native run-ios  # or run-android
```

### 5. Access Applications
- **Frontend**: http://localhost:3004
- **Backend API**: http://localhost:3003
- **API Documentation**: http://localhost:3003/api/v1/health

## 🔧 Environment Configuration

### Backend (.env)
```env
# Server Configuration
PORT=3003
ENV=development

# Supabase Configuration
SUPABASE_URL=postgresql://postgres:[YOUR-PASSWORD]@db.[YOUR-PROJECT-REF].supabase.co:5432/postgres
SUPABASE_ANON_KEY=your_supabase_anon_key
SUPABASE_SERVICE_ROLE_KEY=your_supabase_service_role_key

# JWT Configuration
JWT_SECRET=your_secure_jwt_secret_here
JWT_EXPIRY=24h

# Payment Gateway Configuration
RAZORPAY_KEY_ID=your_razorpay_key_id
RAZORPAY_KEY_SECRET=your_razorpay_key_secret
PAYSTACK_SECRET_KEY=your_paystack_secret_key
PAYSTACK_PUBLIC_KEY=your_paystack_public_key

# Coin System Configuration
WELCOME_COINS=50
MIN_COINS_FOR_PURCHASE=10

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3003,https://yourdomain.com
```

### Frontend (.env)
```env
REACT_APP_API_URL=http://localhost:3003/api/v1
REACT_APP_SUPABASE_URL=your_supabase_url
REACT_APP_SUPABASE_ANON_KEY=your_supabase_anon_key
```

## 📊 Database Schema

The app uses Supabase with the following main tables:
- `users` - User profiles and coin balances
- `series` - Audio series metadata
- `episodes` - Individual episodes with audio files
- `purchases` - User purchase history
- `coin_transactions` - Coin balance changes
- `payments` - Payment gateway transactions
- `coin_bundles` - Available coin packages

## 🔗 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Token refresh

### Series & Episodes
- `GET /api/v1/series` - List all series
- `GET /api/v1/series/:id` - Get series with episodes
- `GET /api/v1/episodes/:id` - Get episode details
- `POST /api/v1/episodes/:id/unlock` - Unlock episode

### User Management
- `GET /api/v1/user/profile` - Get user profile
- `GET /api/v1/user/purchases` - Get purchase history
- `GET /api/v1/user/coins` - Get coin balance

### Payments
- `GET /api/v1/payment/bundles` - Get coin bundles
- `POST /api/v1/payment/initiate` - Start payment
- `POST /api/v1/payment/callback/:gateway` - Payment webhook

## 💰 Payment Integration

- **Razorpay**: Primary payment gateway for India
- **Paystack**: Primary payment gateway for Nigeria
- Coin bundles: 50 coins for ₹50, 120 coins for ₹99, etc.

## 🎯 User Flow

1. User signs up → receives welcome coins
2. Browses series → sees locked episodes
3. Unlocks episodes with coins
4. Can purchase more coins via payment gateways
5. Streams unlocked episodes

## 📱 Mobile Features

- Native audio player with background playback
- Offline episode downloads
- Push notifications for new episodes
- Biometric authentication

## 🌐 Web Features

- Responsive design for desktop and tablet
- Advanced audio player with keyboard shortcuts
- Admin dashboard for content management
- Real-time coin balance updates

## 🔒 Security

- JWT-based authentication
- Role-based access control
- Secure payment processing
- Audio file encryption
- Direct database connection with SSL

## 📈 Scalability

- Microservices-ready architecture
- CDN integration for audio delivery
- Horizontal scaling support
- Caching layer for performance
- Connection pooling with Supabase

## 🛠 Development

### Testing the Connection
```bash
cd backend
go run cmd/server/main.go
```

You should see:
```
✅ Successfully connected to Supabase database
🚀 Server starting on port 3003
```

### Database Operations
The backend now uses real Supabase database operations:
- User registration and authentication
- Series and episode management
- Purchase tracking
- Coin transactions
- Payment processing

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

MIT License - see LICENSE file for details. 