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
- **Database**: Supabase PostgreSQL
- **Authentication**: Supabase Auth
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

## Quick Start

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd audio-series-app
   ```

2. **Set up environment variables**
   ```bash
   # Backend
   cp backend/env.example backend/.env
   # Edit backend/.env with your Supabase credentials
   
   # Frontend
   cp frontend/.env.example frontend/.env
   # Edit frontend/.env with your configuration
   ```

3. **Install dependencies**
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

4. **Set up database**
   - Create a Supabase project
   - Run the SQL schema from `database/schema.sql`
   - Update environment variables with your Supabase credentials

5. **Start the applications**
   ```bash
   # Backend (port 3003)
   cd backend
   go run cmd/server/main.go
   
   # Frontend (port 3004)
   cd frontend
   PORT=3004 npm start
   
   # Mobile
   cd mobile
   npx react-native run-ios  # or run-android
   ```

6. **Access the applications**
   - Frontend: http://localhost:3004
   - Backend API: http://localhost:3003
   - API Documentation: http://localhost:3003/api/v1/health

## 🔧 Environment Configuration

Create `.env` files in each directory with your Supabase and payment gateway credentials.

## 📊 Database Schema

The app uses Supabase with the following main tables:
- `users` - User profiles and coin balances
- `series` - Audio series metadata
- `episodes` - Individual episodes with audio files
- `purchases` - User purchase history
- `coin_transactions` - Coin balance changes

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

## 📈 Scalability

- Microservices-ready architecture
- CDN integration for audio delivery
- Horizontal scaling support
- Caching layer for performance

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

MIT License - see LICENSE file for details. 