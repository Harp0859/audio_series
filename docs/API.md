# Audio Series App API Documentation

## Overview

The Audio Series App API is built with Go and Gin framework, providing endpoints for user authentication, series management, episode streaming, coin system, and payment processing.

## Base URL

```
http://localhost:8081/api/v1
```

## Authentication

Most endpoints require authentication via JWT tokens. Include the token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

## Endpoints

### Authentication

#### POST /auth/register
Register a new user account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123",
  "firstName": "John",
  "lastName": "Doe",
  "phone": "+1234567890"
}
```

**Response:**
```json
{
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "coinBalance": 50,
    "role": "user"
  },
  "token": "jwt-token"
}
```

#### POST /auth/login
Authenticate a user.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "coinBalance": 100,
    "role": "user"
  },
  "token": "jwt-token"
}
```

#### POST /auth/refresh
Refresh the JWT token.

**Headers:** `Authorization: Bearer <token>`

**Response:**
```json
{
  "token": "new-jwt-token"
}
```

### Series

#### GET /series
Get all available series.

**Response:**
```json
[
  {
    "id": "uuid",
    "title": "Forbidden Nights",
    "description": "A thrilling audio series about mystery and suspense",
    "coverImage": "https://example.com/cover1.jpg",
    "author": "Jane Smith",
    "category": "Mystery",
    "isPremium": true,
    "totalEpisodes": 10,
    "createdAt": "2023-01-01T00:00:00Z",
    "updatedAt": "2023-01-01T00:00:00Z"
  }
]
```

#### GET /series/:id
Get a specific series with its episodes.

**Response:**
```json
{
  "series": {
    "id": "uuid",
    "title": "Forbidden Nights",
    "description": "A thrilling audio series about mystery and suspense",
    "coverImage": "https://example.com/cover1.jpg",
    "author": "Jane Smith",
    "category": "Mystery",
    "isPremium": true,
    "totalEpisodes": 10,
    "createdAt": "2023-01-01T00:00:00Z",
    "updatedAt": "2023-01-01T00:00:00Z"
  },
  "episodes": [
    {
      "id": "uuid",
      "seriesId": "uuid",
      "title": "Episode 1: The Beginning",
      "description": "The story begins with a mysterious discovery",
      "audioUrl": "https://example.com/audio1.mp3",
      "duration": 1800,
      "episodeNumber": 1,
      "coinPrice": 10,
      "isLocked": true,
      "createdAt": "2023-01-01T00:00:00Z",
      "updatedAt": "2023-01-01T00:00:00Z"
    }
  ]
}
```

### Episodes

#### GET /episodes/:id
Get episode details with purchase status.

**Headers:** `Authorization: Bearer <token>`

**Response:**
```json
{
  "episode": {
    "id": "uuid",
    "seriesId": "uuid",
    "title": "Episode 1: The Beginning",
    "description": "The story begins with a mysterious discovery",
    "audioUrl": "https://example.com/audio1.mp3",
    "duration": 1800,
    "episodeNumber": 1,
    "coinPrice": 10,
    "isLocked": true,
    "createdAt": "2023-01-01T00:00:00Z",
    "updatedAt": "2023-01-01T00:00:00Z"
  },
  "isOwned": false,
  "canUnlock": true
}
```

#### POST /episodes/:id/unlock
Unlock an episode using coins.

**Headers:** `Authorization: Bearer <token>`

**Response:**
```json
{
  "message": "Episode unlocked successfully"
}
```

#### POST /series/:id/unlock
Unlock an entire series using coins.

**Headers:** `Authorization: Bearer <token>`

**Response:**
```json
{
  "message": "Series unlocked successfully"
}
```

### User

#### GET /user/profile
Get current user's profile.

**Headers:** `Authorization: Bearer <token>`

**Response:**
```json
{
  "id": "uuid",
  "email": "user@example.com",
  "firstName": "John",
  "lastName": "Doe",
  "phone": "+1234567890",
  "avatarUrl": "https://example.com/avatar.jpg",
  "coinBalance": 100,
  "role": "user",
  "isActive": true,
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

#### GET /user/purchases
Get user's purchase history.

**Headers:** `Authorization: Bearer <token>`

**Response:**
```json
[
  {
    "id": "uuid",
    "userId": "uuid",
    "episodeId": "uuid",
    "seriesId": null,
    "type": "episode",
    "amount": 10,
    "paymentId": "payment-ref",
    "status": "completed",
    "createdAt": "2023-01-01T00:00:00Z"
  }
]
```

#### GET /user/coins
Get user's coin balance.

**Headers:** `Authorization: Bearer <token>`

**Response:**
```json
{
  "balance": 100
}
```

### Payments

#### GET /payment/bundles
Get available coin bundles.

**Query Parameters:**
- `currency` (optional): Currency code (INR, NGN). Default: INR

**Response:**
```json
[
  {
    "id": "uuid",
    "name": "50 Coins",
    "coins": 50,
    "price": 5000,
    "currency": "INR",
    "isActive": true,
    "createdAt": "2023-01-01T00:00:00Z"
  }
]
```

#### POST /payment/initiate
Initiate a payment for coin purchase.

**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "bundleId": "uuid",
  "currency": "INR"
}
```

**Response:**
```json
{
  "paymentId": "uuid",
  "gatewayRef": "rzp_1234567890",
  "amount": 5000,
  "currency": "INR",
  "gateway": "razorpay",
  "redirectUrl": "https://checkout.razorpay.com/v1/checkout.html?rzp_1234567890"
}
```

#### POST /payment/callback/:gateway
Handle payment gateway callbacks.

**Request Body:**
```json
{
  "razorpay_payment_id": "pay_1234567890",
  "razorpay_order_id": "order_1234567890",
  "razorpay_signature": "signature"
}
```

**Response:**
```json
{
  "message": "Payment processed successfully"
}
```

### Admin

#### POST /admin/series
Create a new series (Admin only).

**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "title": "New Series",
  "description": "A new audio series",
  "coverImage": "https://example.com/cover.jpg",
  "author": "Author Name",
  "category": "Mystery",
  "isPremium": false
}
```

**Response:**
```json
{
  "id": "uuid",
  "title": "New Series",
  "description": "A new audio series",
  "coverImage": "https://example.com/cover.jpg",
  "author": "Author Name",
  "category": "Mystery",
  "isPremium": false,
  "totalEpisodes": 0,
  "createdBy": "uuid",
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

#### POST /admin/episodes
Create a new episode (Admin only).

**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "seriesId": "uuid",
  "title": "Episode 1",
  "description": "Episode description",
  "audioUrl": "https://example.com/audio.mp3",
  "duration": 1800,
  "episodeNumber": 1,
  "coinPrice": 10,
  "isLocked": true
}
```

**Response:**
```json
{
  "id": "uuid",
  "seriesId": "uuid",
  "title": "Episode 1",
  "description": "Episode description",
  "audioUrl": "https://example.com/audio.mp3",
  "duration": 1800,
  "episodeNumber": 1,
  "coinPrice": 10,
  "isLocked": true,
  "createdAt": "2023-01-01T00:00:00Z",
  "updatedAt": "2023-01-01T00:00:00Z"
}
```

#### GET /admin/stats
Get admin dashboard statistics (Admin only).

**Headers:** `Authorization: Bearer <token>`

**Response:**
```json
{
  "totalUsers": 100,
  "totalSeries": 5,
  "totalEpisodes": 25,
  "totalRevenue": 50000,
  "monthlyRevenue": 10000,
  "activeUsers": 75
}
```

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request
```json
{
  "error": "Invalid request data"
}
```

### 401 Unauthorized
```json
{
  "error": "Authorization header required"
}
```

### 403 Forbidden
```json
{
  "error": "Admin access required"
}
```

### 404 Not Found
```json
{
  "error": "Series not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Failed to get user profile"
}
```

## Coin System

The app uses a virtual coin system where:
- New users receive 50 welcome coins
- Episodes cost 5-20 coins to unlock
- Coins can be purchased via payment gateways
- Coin bundles are available in INR and NGN

## Payment Gateways

### Razorpay (India)
- Supports INR currency
- Payment amounts in paise (smallest unit)
- Webhook integration for payment verification

### Paystack (Nigeria)
- Supports NGN currency
- Payment amounts in kobo (smallest unit)
- Webhook integration for payment verification

## Rate Limiting

API endpoints are rate-limited to prevent abuse:
- 100 requests per minute per IP
- 1000 requests per hour per user

## CORS

The API supports CORS for cross-origin requests:
- Allowed origins: http://localhost:3000, http://localhost:3001
- Allowed methods: GET, POST, PUT, DELETE, OPTIONS
- Allowed headers: Content-Type, Authorization 