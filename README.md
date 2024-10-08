# Hexagonal API v2

This is the second version of the API, designed using the hexagonal architecture pattern. It handles operations related to students ("alumno") and teachers ("maestro"). The architecture ensures scalability, maintainability, and flexibility by decoupling core business logic from external dependencies such as databases and frameworks.

## Table of Contents

- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
  - [Running the API](#running-the-api)
  - [Database Migrations](#database-migrations)
- [Environment Variables](#environment-variables)
- [Dependencies](#dependencies)
- [API Endpoints](#api-endpoints)
- [License](#license)

## Project Structure

```plaintext
src
│
├── alumno
│   ├── application
│   ├── domain
│   └── infraestructure
├── maestro
│   ├── application
│   ├── domain
│   └── infraestructure
│
├── bootstrap
│   └── bootstrap.go            
│
└── shared
    ├── database
    │   ├── database.go
    │   └── migrations
    │       ├── 20240813_create_alumno_table.up.sql
    │       └── 20240813_create_maestro_table.up.sql
    ├── middleware
    │   └── jwt_middleware.go   
    └── auth
        └── auth.go             
│
└── main.go
```

## Installation

To run this project locally, follow these steps:

1. **Clone the repository:**

   ```
   git clone https://github.com/Lalo64GG/GO-HEXAGONAL-V2.git
   cd GO-HEXAGONAL-V2
   ```
2. **Install dependencies:**
   ```bash
   go mod tidy
   ```
## Usage

 ### Running the API
   ```
   go mod tidy
   ```
   This will launch the API on the default port (usually :8080). You can configure the port by setting the appropriate environment variable.

 ### Database Migrations
 Apply database migrations to set up your database schema:
  ```
   migrate -path src/shared/database/migrations -database "mysql://your_user:yourpassword@tcp(127.0.0.1:3306)/your_dbName?parseTime=true" up

   ```
   
## Dependencies
  The project uses the following major dependencies:

-   Gin: A web framework for Go, used for handling HTTP requests.
-   godotenv: Loads environment variables from a .env file.
-   golang-migrate: A tool for managing database migrations.
-   JWT: Used for secure authentication in the middleware.
-   Install them with go mod tidy after cloning the repository.

## API Endpoints
Here are some of the main API endpoints:

- GET /alumno: Retrieve a list of students.
- POST /alumno: Create a new student.
- GET /maestro: Retrieve a list of teachers.
- POST /maestro: Create a new teacher.
- Documentation for all endpoints will be provided separately.