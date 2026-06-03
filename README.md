# Weather Dashboard

A full-stack weather application built with Go and Next.js. Search for weather by city or use your current location, view 5-day forecasts, and save up to 3 favourite cities.

---

## Tech Stack

**Backend:** Go, Fiber, pgxpool, Redis, JWT  
**Frontend:** Next.js (App Router), TypeScript, TanStack Query, Tailwind, Sonner  
**Database:** PostgreSQL  
**API:** OpenWeatherMap

---

## Prerequisites

- Go 1.21+
- Node.js 18+
- PostgreSQL
- Redis
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- OpenWeatherMap API key (free at openweathermap.org)

## Backend Setup

### 1. Environment Variables

Copy `.env.example` to `.env` inside the `backend/` folder:

```bash
cp backend/.env.example backend/.env
```

Fill in the values:

```env
PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/weather_app
TEST_DATABASE_URL=postgres://user:password@localhost:5432/weather_app_test
REDIS_URL=redis://localhost:6379
TEST_REDIS_URL=redis://localhost:6379
JWT_SECRET=your_super_secret_key_here
OPENWEATHER_API_KEY=your_openweather_key_here
CLIENT_URL=http://localhost:3000
```

### 2. Run Migrations

```bash
migrate -path ./backend/migrations \
  -database "postgres://user:password@localhost:5432/weather_app?sslmode=disable" up
```

### 3. Start the Backend

```bash
cd backend
go run cmd/main.go
```

Server runs on `http://localhost:8080`

### 4. Run Tests

```bash
cd backend
go test ./tests/... -v
```

---

## Frontend Setup

### 1. Install Dependencies

```bash
cd frontend
npm install
```

### 2. Start the Frontend

```bash
npm run dev
```

App runs on `http://localhost:3000`

---

## API Endpoints

### Auth

| Method | Endpoint             | Description           |
| ------ | -------------------- | --------------------- |
| POST   | `/api/auth/register` | Register a new user   |
| POST   | `/api/auth/login`    | Login and receive JWT |

### Weather (Protected)

| Method | Endpoint                           | Description                    |
| ------ | ---------------------------------- | ------------------------------ |
| GET    | `/api/weather/:city`               | Current weather by city        |
| GET    | `/api/weather/coords?lat=x&lon=y`  | Current weather by coordinates |
| GET    | `/api/forecast/:city`              | 5-day forecast by city         |
| GET    | `/api/forecast/coords?lat=x&lon=y` | 5-day forecast by coordinates  |

### Favorites (Protected)

| Method | Endpoint               | Description                       |
| ------ | ---------------------- | --------------------------------- |
| GET    | `/api/favorites`       | Get saved favourites with weather |
| POST   | `/api/favorites`       | Add a city to favourites (max 3)  |
| DELETE | `/api/favorites/:city` | Remove a city from favourites     |

All protected endpoints require `Authorization: Bearer <token>` header.

---

## Features

- JWT authentication stored in cookies
- Weather search by city name
- Geolocation support — use your current location
- 5-day forecast with daily high/low temperatures
- Save up to 3 favourite cities
- Redis caching — 10 min for current weather, 30 min for forecasts
- Rate limiting on auth endpoints
- Skeleton loading states
- Toast notifications

---

## Postman Collection

Import `weather-app.postman_collection.json` from the repo root into Postman. Set `BASE_URL_WEATHER_APP` to `http://localhost:8080/api` and run Login to auto-populate the `token` variable.
