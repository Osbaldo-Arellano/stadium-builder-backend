### Stadium Builder Backend

Stadium Builder Backend is the server-side component of the **Stadium Builder** game, where players can bet on real-world sports events, use their winnings to base-build, and track their progress through leaderboards. This backend integrates with external sports betting APIs, manages caching using Redis, and provides RESTful API endpoints for interacting with the game data.

---

### **Features**

1. **Betting Data Integration**

   - Fetches real-world sports betting data (e.g., OddsAPI).
   - Periodically updates betting data using a scheduler.
   - Implements caching for efficient and fast data retrieval.

2. **Leaderboard Service**

   - Stores and retrieves player rankings.
   - Provides CRUD operations for leaderboard management.
   - Ensures persistence using PostgreSQL.

3. **Game Data Models**

   - Players: Stores player progress and data.
   - Bets: Manages real-time betting data.
   - Stadiums: Customization options for player bases.
   - Leaderboards: Tracks global rankings.

4. **API Endpoints**

   - Retrieve betting odds.
   - Fetch and update leaderboard scores.
   - Health checks for service availability.

5. **Efficient Caching**
   - Uses Redis to cache betting data, reducing API call frequency.
   - Ensures consistent and updated data availability for clients.

---

### **Technologies Used**

- **Programming Language:** Go
- **Frameworks:**
  - Gin (HTTP routing)
  - Go-Coop Gocron (Scheduler)
- **Databases:**
  - PostgreSQL (Relational database)
  - Redis (Caching)
- **External API:** OddsAPI for real-world sports data.
- **Testing Frameworks:**
  - Testify for unit testing
  - Mocking with `httptest`

---

### **Project Structure**

```plaintext
stadium-builder-backend/
├── config/            # Configuration for PostgreSQL and Redis
├── models/            # Data models for Players, Bets, Leaderboards, etc.
├── routes/            # API routes
├── services/          # Core business logic (e.g., Fetching data, caching)
├── unit_tests/        # Unit tests for the project
├── main.go            # Application entry point
├── go.mod             # Module dependencies
├── go.sum             # Dependency checksums
```

---

### **Setup Instructions**

#### **1. Prerequisites**

- Go 1.21 or later
- Docker (optional for running dependencies locally)
- PostgreSQL
- Redis

#### **2. Clone the Repository**

```bash
git clone https://github.com/your-username/stadium-builder-backend.git
cd stadium-builder-backend
```

#### **3. Set Up Environment Variables**

Create a `.env` file with the following variables:

```plaintext
PORT=8080
DATABASE_URL=postgres://<username>:<password>@localhost:5432/<database_name>
REDIS_URL=localhost:6379
BETTING_API_KEY=<your_api_key>
ODDS_API_URL=https://api.the-odds-api.com/v4/sports/americanfootball_nfl/odds
```

#### **4. Run PostgreSQL and Redis**

If using Docker:

```bash
docker run -d -p 5432:5432 --name postgres -e POSTGRES_USER=stadium_user -e POSTGRES_PASSWORD=secure_password -e POSTGRES_DB=stadium_builder postgres
docker run -d -p 6379:6379 --name redis redis
```

#### **5. Build and Run the Project**

```bash
go build -o stadium-backend .
./stadium-backend
```

---

### **API Endpoints**

#### **Health Check**

- `GET /health`
  - **Description:** Check service availability.
  - **Response:** `{"status": "healthy"}`

#### **Fetch Betting Data**

- `GET /betting`
  - **Description:** Retrieve real-world sports betting odds.
  - **Response:** JSON list of games and odds.

#### **Leaderboard Management**

- `GET /leaderboard`
  - **Description:** Fetch the top leaderboard entries.
- `POST /leaderboard`
  - **Description:** Update or add a leaderboard entry.

---

### **Testing**

Run unit tests:

```bash
go test ./unit_tests/... -v
```

---

### **CI/CD Pipeline**

The project uses GitHub Actions for Continuous Integration.

#### Workflow Features:

- Spins up PostgreSQL and Redis services.
- Runs unit tests to ensure code reliability.

---

### **Future Enhancements**

- **Player Authentication:** Add secure login and account management.
- **Advanced Betting Features:** Include more sports and betting options.
- **Real-time Notifications:** Notify players of betting results and leaderboard changes.

---
