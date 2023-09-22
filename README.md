# Description

project golang fiber (api) connect postgres database with loadtest locust

## Postgres

init table

- product

init data to this table 30 row.

## How to check init database

1. connect to postgres containner.

```bash
docker exec -it postgres /bin/bash
```

2.connect to postgres.

```bash
psql -U postgres
```

3. use this commabd to list database table.

```bash
\dt
```

**if you seen table product -> it OK.**

## How To Run

```bash
docker compose up
```

when database not init use this

```bash
docker compose down --volumes
```

## Detail

- postgres -> port 5432
- golang -> port 3000
- locust -> port 8089
