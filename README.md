# Warehouse Management System

## Authors
**Adil Ormanov & Madi Kassymov**
**Group**: SE-2311

## Project Description
The Warehouse Management System is a web application designed to manage inventory for warehouses.
- **Administrators**: Can add, update, delete, and view products.
- **Cashiers**: Manage product quantities.
- **Unregistered Users**: Can only view available products.

## Homepage Screenshot
![Homepage Screenshot](warehouse-frontend/images/Screenshot%202024-12-20%20at%2007.19.04.png)

## Project Diagram
For an overview of the system design, refer to the online diagram:
[View Project Diagram](https://drive.google.com/file/d/1kKt8_p4BTikiviT1lTGdPM7cTFXPPR47/view?usp=sharing)

---

## Used Tools
### Development Tools
- **Programming Language**: [Go (Golang)](https://golang.org/)
- **Frontend Technologies**: HTML, CSS, JavaScript
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **Version Control**: [Git](https://git-scm.com/)
- **IDE**: JetBrains GoLand
- **Other Tools**: Google Drive (for diagrams and collaboration)

### Libraries and Frameworks
#### Backend Dependencies (Go Modules)
Here is the list of Go libraries currently used in the project (from `go.mod`):
- `github.com/dgrijalva/jwt-go` - JWT implementation for authentication.
- `github.com/go-chi/chi` - Lightweight and idiomatic Go web framework.
- `github.com/lib/pq` - PostgreSQL driver for Go.
- `gorm.io/gorm` - ORM library for Go to work with relational databases.
- `gorm.io/driver/postgres` - PostgreSQL driver package for GORM.

#### Frontend Dependencies
- **No package manager**: No NPM or third-party libraries are yet used for the frontend—uses plain HTML, CSS, and
  JavaScript.

---

## Setup Instructions

### Clone the Repository
1. Clone the repository and navigate to the backend folder:
```bash
git clone https://github.com/Adilforest/Assignment_1_AP.git
cd warehouse-backend
```

### Backend Setup
1. **Install Go**:

- Download and install Go from the [official website](https://golang.org/dl/).
- Verify the installation:
```bash
go version
```
2. **Install PostgreSQL**:

- Download PostgreSQL from [PostgreSQL.org](https://www.postgresql.org/download/).
- Verify the installation:
```bash
psql --version
```
3. **Configure the Database**:

- Create a PostgreSQL database and user with these SQL commands:
    ```sql
    CREATE DATABASE warehouse;
    CREATE USER admin WITH PASSWORD '<yourpassword>';
    ALTER ROLE admin SET client_encoding TO 'utf8';
    ALTER ROLE admin SET default_transaction_isolation TO 'read committed';
    ALTER ROLE admin SET timezone TO 'UTC';
    GRANT ALL PRIVILEGES ON DATABASE warehouse TO admin;
    ```

    4. **Install Project Dependencies**:

    - Navigate to the backend folder:
    ```bash
    cd warehouse-backend
    ```
    - Install dependencies:
    ```bash
    go mod tidy
    ```

    5. **Run the Server**:

    - Start the development server:
    ```bash
    go run main.go
    ```
    - The backend should now be running at: [http://localhost:8080](http://localhost:8080)

  ### Frontend Setup
    1. **Navigate to the Frontend Folder**:
    ```bash
    cd warehouse-frontend
    ```
    2. **Open the Frontend**:

    - Open `index.html` in your browser:
    ```bash
    start index.html
    ```

    ---

  ## File Structure Overview

  ### Backend Directory
    ```plaintext
    /warehouse-backend
    ├── main.go                     // Entry point for server initialization.
    ├── /routes                     // Route definitions.
    │   ├── auth_routes.go          // Authentication routes.
    │   ├── product_routes.go       // Product management routes.
    ├── /controllers                // Business logic layer.
    │   ├── auth_controller.go      // Authentication logic.
    │   ├── product_controller.go   // Product logic.
    ├── /models                     // Database schema definitions.
    │   ├── user.go                 // User schema.
    │   ├── product.go              // Product schema.
    ├── /middleware                 // Middleware for request handling.
    │   ├── auth_middleware.go      // JWT authentication middleware.
    │   ├── role_middleware.go      // Role-based access control.
    ├── /database                   // Database connection logic.
    │   ├── mongo_connection.go     // MongoDB connection (placeholder).
    │   ├── postgres_connection.go  // PostgreSQL connection.
    └── /config                     // Application configuration.
    ├── config.go
    ```

  ### Frontend Directory
    ```plaintext
    /warehouse-frontend
    ├── index.html                  // Guest access to view products.
    ├── admin.html                  // Admin dashboard.
    ├── cashier.html                // Cashier interface.
    ├── /js
    │   ├── auth.js                 // Login/logout logic.
    │   ├── products.js             // API calls for product operations.
    │   ├── api.js                  // Centralized API request handling.
    ├── /css
    │   ├── styles.css              // Application styles.
    ├── /images                     // Contains images used in the
    ```

---