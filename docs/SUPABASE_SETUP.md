# 🗄️ Supabase Setup Guide

This guide will help you set up Supabase as the backend database for the Audio Series App using direct database connection.

## 📋 Prerequisites

- A Supabase account (free tier available)
- Go 1.21+ installed
- Basic knowledge of PostgreSQL

## 🚀 Step-by-Step Setup

### 1. Create a Supabase Project

1. Go to [Supabase Dashboard](https://supabase.com/dashboard)
2. Click "New Project"
3. Choose your organization
4. Enter project details:
   - **Name**: `audio-series-app` (or your preferred name)
   - **Database Password**: Choose a strong password
   - **Region**: Select the closest region to your users
5. Click "Create new project"
6. Wait for the project to be ready (usually 1-2 minutes)

### 2. Get Database Connection String

1. In your Supabase project dashboard, go to **Settings > Database**
2. Scroll down to the **Connection string** section
3. Copy the **Direct connection** string that looks like:
   ```
   postgresql://postgres:[YOUR-PASSWORD]@db.[YOUR-PROJECT-REF].supabase.co:5432/postgres
   ```

### 3. Get API Keys

1. In your Supabase project dashboard, go to **Settings > API**
2. Copy the following keys:
   - **anon public**: Your project's anon key
   - **service_role secret**: Your project's service role key

### 4. Configure Environment Variables

1. Copy the environment template:
   ```bash
   cp backend/env.example backend/.env
   ```

2. Edit `backend/.env` with your Supabase credentials:
   ```env
   # Supabase Configuration
   SUPABASE_URL=postgresql://postgres:[YOUR-PASSWORD]@db.[YOUR-PROJECT-REF].supabase.co:5432/postgres
   SUPABASE_ANON_KEY=your_supabase_anon_key
   SUPABASE_SERVICE_ROLE_KEY=your_supabase_service_role_key
   
   # JWT Configuration
   JWT_SECRET=your_secure_jwt_secret_here
   JWT_EXPIRY=24h
   ```

### 5. Setup Database Schema

1. Go to your Supabase project's **SQL Editor**
2. Copy the entire contents of `database/schema.sql`
3. Paste it into the SQL Editor
4. Click "Run" to execute the schema
5. Verify the tables are created in **Table Editor**

### 6. Test the Connection

Run the connection test:
```bash
cd backend
go run test_connection.go
```

You should see:
```
🔗 Testing Supabase connection...
URL: postgresql://postgres:***@db.***.supabase.co:5432/postgres
✅ Successfully connected to Supabase database!
🎉 Connection test completed successfully!
```

### 7. Start the Backend Server

```bash
cd backend
go run cmd/server/main.go
```

You should see:
```
✅ Successfully connected to Supabase database
🚀 Server starting on port 3003
```

## 🔧 Configuration Details

### Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `SUPABASE_URL` | Direct database connection string | `postgresql://postgres:password@db.ref.supabase.co:5432/postgres` |
| `SUPABASE_ANON_KEY` | Public API key for client access | `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...` |
| `SUPABASE_SERVICE_ROLE_KEY` | Secret key for admin operations | `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...` |
| `JWT_SECRET` | Secret for JWT token signing | `your-secure-secret-here` |

### Database Schema

The application uses the following tables:

- **users**: User profiles and coin balances
- **series**: Audio series metadata
- **episodes**: Individual episodes with audio files
- **purchases**: User purchase history
- **coin_transactions**: Coin balance changes
- **payments**: Payment gateway transactions
- **coin_bundles**: Available coin packages

## 🔒 Security Considerations

### Connection Security
- Uses SSL/TLS encryption by default
- Direct connection bypasses Supabase's API layer
- Requires proper firewall configuration

### API Keys
- **anon key**: Safe for client-side use
- **service_role key**: Keep secret, only for server-side operations
- Rotate keys regularly in production

### JWT Configuration
- Use a strong, random JWT secret
- Set appropriate expiry times
- Consider using environment-specific secrets

## 🚨 Troubleshooting

### Connection Issues

**Error**: `Failed to connect to database`
- Check your `SUPABASE_URL` format
- Verify your database password
- Ensure your IP is not blocked by firewall

**Error**: `Failed to ping database`
- Check network connectivity
- Verify Supabase project is active
- Check if database is in maintenance mode

### Schema Issues

**Error**: `relation "users" does not exist`
- Run the schema SQL in Supabase SQL Editor
- Check for SQL syntax errors
- Verify all tables are created

### Authentication Issues

**Error**: `JWT_SECRET is required`
- Set a secure JWT secret in your `.env` file
- Use a random string of at least 32 characters

## 📊 Monitoring

### Database Metrics
- Monitor connection pool usage
- Track query performance
- Set up alerts for connection failures

### Application Logs
- Check server logs for database errors
- Monitor API response times
- Set up error tracking

## 🔄 Migration Strategy

### Development to Production
1. Create separate Supabase projects for dev/staging/prod
2. Use environment-specific connection strings
3. Test migrations on staging first
4. Backup production data before major changes

### Schema Updates
1. Create migration scripts
2. Test on staging environment
3. Schedule maintenance windows
4. Monitor for data integrity issues

## 📚 Additional Resources

- [Supabase Documentation](https://supabase.com/docs)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Go Database/SQL Tutorial](https://golang.org/doc/database)
- [JWT Best Practices](https://auth0.com/blog/a-look-at-the-latest-draft-for-jwt-bcp/)

## 🆘 Support

If you encounter issues:

1. Check the [Supabase Status Page](https://status.supabase.com/)
2. Review the troubleshooting section above
3. Check the application logs for specific error messages
4. Verify your configuration matches the examples

---

**🎧 Audio Series App** - Complete Supabase integration with direct database connection for optimal performance and security. 