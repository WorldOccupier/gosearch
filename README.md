# gosearch

Backend for retrieving search results using TF-IDF ranking.

## Architecture

- **Framework:** Gin (HTTP router)
- **Database:** PostgreSQL with pgx/v5 driver
- **Search:** TF-IDF scoring via SQL query over `t_url_term_count` and `t_metadata` tables

## Prerequisites

- Go 1.26+
- Docker & Docker Compose (for containerized run)
- PostgreSQL (for local-only run)

## Quick Start (Docker)

```bash
make up
```

This builds the Docker image and starts the service on port `8080`.

You can also run directly:

```bash
docker compose up -d
```

## Local Build (without Docker)

```bash
make localbuild
```

Then run the binary with a `DATABASE_URL` environment variable pointing to your Postgres instance:

```bash
DATABASE_URL="postgres://user:pass@localhost:5432/goprocess" ./gosearch
```

## API

**GET /search?query=<terms>**

Returns top 10 results ranked by TF-IDF score. Example:

```bash
curl "http://localhost:8080/search?query=golang"
```

## Makefile Targets

| Target       | Description                          |
|-------------|--------------------------------------|
| `up`        | Build and start containers           |
| `down`      | Stop containers                      |
| `rebuild`   | Stop, remove volumes, rebuild, start |
| `logs`      | Follow container logs                |
| `clean`     | Stop containers and remove volumes   |
| `localbuild`| Build binary locally                 |

## Cleanup

```bash
make clean
```

Or manually:

```bash
docker compose down -v
```

This stops containers and removes persisted volumes.
