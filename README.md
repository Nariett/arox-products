# Arox-products

This `.env` file is used to configure the PostgreSQL database connection and connect with another services for the **Arox** project.

## Environment Variables

| Variable     | Description                         | Example Value       |
|--------------|-------------------------------------|---------------------|
| `DB_USER`    | PostgreSQL username                 | `postgres`          |
| `DB_PASSWORD` | Password for the specified user     | `postgres`          |
| `DB_NAME`    | Name of the database                | `arox-gateway`      |
| `DB_SSLMODE` | SSL connection mode                 | `disable` (default) |
| `DB_PORT`    | Port on which PostgreSQL is running | `5432`              |
| `HOST`       | Host where the database is located  | `localhost`         |
| `LPORT`      | Port to connect with arox-gateway   | `8001`              |
| `PROTOCOL`   | Protocol user for listen GRPC port  | `tcp`     

## Example Connection String

```text
postgres://postgres:postgres@localhost:5432/arox-gateway?sslmode=disable
