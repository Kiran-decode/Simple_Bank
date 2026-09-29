Simple Bank

A backend banking application built using Go, Gin, PostgreSQL, SQLC, Docker, and database migrations.

The project provides REST APIs for managing bank accounts and performing banking operations such as creating accounts, retrieving account information, and transferring money between accounts.

🚀 Tech Stack

- Go — Backend programming language
- Gin — HTTP web framework
- PostgreSQL — Relational database
- SQLC — Type-safe Go code generation from SQL queries
- Docker — PostgreSQL containerization
- golang-migrate — Database migrations
- GitHub Actions — Continuous Integration
- Make — Development automation

📌 Features

- Create bank accounts
- Retrieve account details
- List accounts
- Transfer money between accounts
- PostgreSQL database integration
- Database migrations
- Type-safe database access using SQLC
- REST API development using Gin
- Unit and integration testing
- Docker-based PostgreSQL setup
- Automated testing with GitHub Actions

🏗️ Project Structure

simplebank/
│
├── api/
│   └── ...                  # API handlers and request/response models
│
├── db/
│   ├── migration/           # Database migration files
│   ├── query/               # SQL queries
│   └── sqlc/                # SQLC generated code
│
├── util/
│   └── ...                  # Configuration and utility functions
│
├── main.go                  # Application entry point
├── server.go                # HTTP server setup
├── Dockerfile
├── docker-compose.yaml
├── Makefile
├── sqlc.yaml
└── go.mod

⚙️ Prerequisites

Make sure you have the following installed:

- Go
- Docker
- PostgreSQL
- Make
- golang-migrate
- SQLC

Check your Go installation:

go version

Check Docker:

docker --version

🐘 PostgreSQL Setup

Start a PostgreSQL container:

docker run --name postgres12 \
  -e POSTGRES_USER=root \
  -e POSTGRES_PASSWORD=secret \
  -p 5432:5432 \
  -d postgres:12-alpine

Check that the container is running:

docker ps

🗄️ Database Setup

Create the database:

make createdb

Run migrations:

make migrateup

To roll back the migrations:

make migratedown

🔧 Configuration

Configure the database connection according to your local environment.

Example:

postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

«Do not commit real database credentials or ".env" files to the repository.»

▶️ Run the Application

Download dependencies:

go mod download

Run the server:

go run main.go

The API will be available at:

http://localhost:8080

🔌 API Endpoints

Accounts

Create an account:

POST /accounts

Get an account:

GET /accounts/:id

List accounts:

GET /accounts

Transfers

Create a money transfer:

POST /transfers

Example request:

{
  "from_account_id": 1,
  "to_account_id": 2,
  "amount": 100
}

🧬 SQLC

This project uses SQLC to generate type-safe Go code from SQL queries.

After modifying the SQL queries, regenerate the Go code:

make sqlc

Using SQLC provides compile-time type safety while still allowing the application to use SQL directly.

🔄 Database Migrations

Database schema changes are managed using migrations.

Create a migration:

make new_migration name=<migration_name>

Example:

make new_migration name=add_users

Apply migrations:

make migrateup

Rollback migrations:

make migratedown

🧪 Testing

Run all tests:

go test ./...

Run tests with coverage:

go test ./... -cover

Generate a coverage report:

go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

🐳 Docker

Docker is used to run PostgreSQL locally without requiring a manual PostgreSQL installation.

The project can also be containerized for a more consistent development environment.

🔄 Continuous Integration

GitHub Actions is used to automate project verification.

The CI workflow can:

1. Set up PostgreSQL
2. Run database migrations
3. Run Go tests
4. Generate test coverage
5. Build the application

📚 What I Learned

Through this project, I have been learning practical backend development concepts including:

- Building REST APIs with Go
- Gin routing and request handling
- PostgreSQL database design
- Writing SQL queries
- SQLC code generation
- Database transactions
- Database migrations
- Unit and integration testing
- Mocking dependencies
- Docker-based development
- Makefile automation
- GitHub Actions and CI
- Structuring a Go backend project

🛠️ Future Improvements

Planned improvements include:

- Add JWT-based authentication
- Add user registration and login
- Improve API documentation with Swagger/OpenAPI
- Add more comprehensive integration tests
- Add structured logging
- Add monitoring and metrics
- Improve CI/CD pipeline
- Deploy the application to a cloud platform

👨‍💻 Author

Kiran PB

B.Tech — Information Technology
NITK Surathkal

Currently learning backend development with Go and Data Structures & Algorithms.

---

⭐ If you find this project useful, feel free to explore the repository.
