# 🔥 IgnisGo Golang Boilerplate 🏗️

Ignis is a production-ready **Go project template** designed to jumpstart your backend development. It provides a solid foundation with best practices, pre-configured tooling, and essential components, acting as a "batteries-included" starter kit for building robust RESTful services.

Think of it as a modern, opinionated copy-paste-able base for your next Go project.

## 🚀 Included Batteries

-   **Project Structure**: Follows Standard Go Project Layout.
-   **PostgreSQL Ready**: Pre-configured `pgx` driver and connection pooling.
-   **Type-Safe SQL**: `sqlc` setup for generating Go code from SQL queries.
-   **Dockerized**: Complete `Dockerfile` and `docker-compose.yml` for local development.
-   **Migrations**: Database migration support via `goose`.
-   **Configuration**: Simple env-var based config with `.env` file support.
-   **Graceful Shutdown**: Ready-to-use signal handling for safe termination.
-   **Health Checks**: Built-in `/healthz` endpoint.
-   **Example Domain**: Includes a minimal `User` domain to demonstrate patterns (CRUD, Repository).

> [!NOTE]
> The included `users` table migration (`00001_initial_schema.sql`) and queries are for **demonstration purposes only**. You should replace them with your own domain logic.

## 🛠️ Tech Stack & Dependencies

| Component | Technology | Version | Role | Source |
| :--- | :--- | :--- | :--- | :--- |
| **Language** | [Go](https://go.dev/) | 1.25.4 | Core programming language | [Official Site](https://go.dev/) |
| **Database** | [PostgreSQL](https://www.postgresql.org/) | 18.1 (Alpine) | Relational database system | [Official Site](https://www.postgresql.org/) |
| **Driver** | [pgx](https://github.com/jackc/pgx) | v5.8.0 | PostgreSQL driver and toolkit | [GitHub](https://github.com/jackc/pgx) |
| **ORM/Gen** | [sqlc](https://sqlc.dev/) | v1.30.0 | Type-safe SQL compiler | [GitHub](https://github.com/sqlc-dev/sqlc) |
| **Config** | [godotenv](https://github.com/joho/godotenv) | v1.5.1 | Environment variable loader | [GitHub](https://github.com/joho/godotenv) |
| **Migration** | [Goose](https://pressly.github.io/goose/) | v3.26.0 | Database migration tool | [GitHub](https://github.com/pressly/goose) |

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

-   **Go**: Version 1.25.4 or higher
-   **Docker**: For containerized execution
-   **Make**: For running project commands

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/ioanzicu/ignis.git
cd ignis
```

### 2. Configuration

The application uses environment variables for configuration.

**Local Development**: Create a `.env` file in the root directory:

```env
HTTP_HOST=localhost
HTTP_PORT=8080
PG_DSN=postgres://user:password@localhost:5432/dbname?sslmode=disable
```

**Docker**: Variables are automatically handled in `docker-compose.yml`.

### 3. Running with Docker (Recommended)

Start the application and database with a single command:

```bash
make compose-up
```

This will: 
- Start Postgres 18.1
- Build and start the `app` service
- Expose the API on port `8080`

To stop the services:
```bash
make compose-down
```

To view logs:
```bash
make compose-logs
```

### 4. Running Locally

Install dependencies and tools:
```bash
make deps
```

Generate database code (if queries change):
```bash
make generate
```

Run the application:
```bash
make run
```

## 🛠️ Make Commands

The project includes a `Makefile` to simplify common tasks:

| Command | Description |
| :--- | :--- |
| `make help` | Display available commands |
| `make deps` | Install Go dependencies and tools (`sqlc`, `goose`) |
| `make run` | Run the application locally |
| `make build` | Build the application binary |
| `make test` | Run unit tests |
| `make generate` | Generate Go code from SQL using `sqlc` |
| `make migration-create` | Create a new database migration |
| `make migration-up` | Apply database migrations |
| `make compose-up` | Start Docker environment |
| `make compose-down` | Stop Docker environment |

## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

---

**Happy Coding!** 🐹 built with ❤️ in Go.