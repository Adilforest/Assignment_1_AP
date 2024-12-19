# Assignment_1_AP
# Adil Ormanov & Madi Kassymov
**Group**: SE-2311

# Warehouse Backend

Warehouse Backend is a lightweight backend service built with Go and the Gin framework. It is designed to handle basic API operations and database interactions for warehouse management.

## Features

- **RESTful API**: Supports GET and POST requests.
- **CORS Support**: Cross-Origin Resource Sharing is enabled for easier development.
- **Database Integration**:
    - PostgreSQL: Primary database with automatic migration capabilities.
    - MongoDB: Placeholder for additional database support.
- **Simple Setup**: Easily configurable and extensible.

## Requirements

- Go 1.20+
- PostgreSQL 16+
- MongoDB 6+

## Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Adilforest/Assignment_1_AP.git
   cd warehouse-backend
   ```

2. **Configure the environment:**
   Set up your configuration file (e.g., `config.yml`) with the necessary database credentials.

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Run the application:**
   ```bash
   go run main.go
   ```

## File Structure

- `main.go`: Entry point of the application.
- `database/postgres_connection.go`: Handles PostgreSQL connection and migrations.
- `database/mongo_connection.go`: Placeholder for MongoDB integration.

## API Endpoints

### Base URL

`http://localhost:8080`

### Endpoints

- `GET /` - Welcome message.
- `GET /get` - Returns a success message.
- `POST /post` - Accepts a JSON payload and returns a success response.

Example payload for `POST /post`:
```json
{
  "message": "Your custom message"
}
```

## Database Setup

### PostgreSQL

1. Ensure PostgreSQL is running.
2. Configure the `PostgresDSN` in your configuration file.
3. The application automatically creates the required tables upon startup.

### MongoDB 

MongoDB integration is currently a placeholder and requires additional implementation.





