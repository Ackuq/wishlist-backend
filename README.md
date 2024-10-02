# Wishlist backend

This is a simple Go backend I created with the purpose of deepning my knowledge of Go.

## Tools used

- [Docker](https://www.docker.com) - Containerization for local development
- [PostgreSQL](https://www.postgresql.org) - Main database used
- [pgx](https://github.com/jackc/pgx) - Go driver for PostgreSQL
- [sqlc](https://github.com/sqlc-dev/sqlc) - SQL Code generation
- [migrate](https://github.com/golang-migrate/migrate) - Database migrations
- [godotenv](https://github.com/joho/godotenv) - .env file management

## Set up

### Environment variables

The file `.env.example` contains all of the required environment variables used to be able to run this project.

Copy the file `.env.example` to a file `.env`. And fill them out as necessary.

```sh
cp .env.example .env
```

## Running (locally)

Start the local database by running `docker-compose`.

```sh
docker-compose up -d
```

When the database is up and running, start the application by running the following Make target.

```sh
make run
```

Your server should now be up and running!

## Migrations

When running migrations, make sure you have started the local database use `docker-compose`, otherwise the following commands will fail.

### Creating new migration

```sh
make create-migration name='create_account_table'
```

Fill in the generated migration files with the required SQL to perform and undo an migration, in the files `[...].up.sql`, and `[...].down.sql` respectfully.

### Migrating your database

Migrating the database is either done when running the application, or by running the following Make target.

```sh
# Run all migrations
make migrate
# Run the next 2 migrations
make migrate n=2
```

### Reverting a migration

Reverting a migration a migration is simple, you can either revert all or the last `n` migrations.

```sh
# Run all the down migrations
make migrate-down
# Run the last 2 down migrations
make migrate-down n=2
```

## Queries

This project uses sqlc to generate Go code bindings from raw SQL queries. To create SQL queries, add the desired SQL to file in `./internal/db/queries/`. Prefix the filename with the table that you are operating on.

Each query should have a comment which specifies the desired Go function name of query, and if it return one or multiple rows. E.g.

```sh
-- name: GetAccount :one
SELECT * FROM account
    WHERE id = $1 LIMIT 1;
```

For full documentation regarding sqlc, [click here](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html).

Generating the code binding is done by running the following command.

```sh
sqlc generate
```
