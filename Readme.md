# Project DataReport (GO)

## Hypothetical case
*The story below is pure fiction*

Imagine that at some point in the editorial office of a newspaper with few resources, articles were stored in SQLite. After modernization, this information was migrated to PostgreSQL.

Well, the work has been done and there are now several ways to validate this data import. A GO application was developed that accessed the two databases and returned the records for comparison.

### Future
For now, the application is only providing a general query endpoint. For the few records that exist in the database, this is enough. But in the future, with thousands records, more endpoints to metrics will be needed...

## How does it work?
The app using two databases, **SQLite** or **PortgreSQL**, to change database engine change **./src/config/app.json**.
- Values valid to index **dbengine**: `sqlite` or `postgres`

### Start
- `git clone https://github.com/psaraiva/data-reporter.git`
- `cd data-reporter`
- `chmod -R 777 ./db`
- Execute step Config *Docker DataBases*
- `cd src`
- `go get -u`
- `go mod tidy`
- `go run main.go`

### Configuration Docker DataBases
Warning: All commands should **executed** in the **root folder project**.

#### PostgreSQL
- `docker-compose -f Docker/postgres/docker-compose.yml up -d`
- `docker cp db/db_script/postgres.sql datareporter_postgres:/docker-entrypoint-initdb.d/dump.sql`
- `docker exec -it datareporter_postgres psql -h localhost -U postgres -f docker-entrypoint-initdb.d/dump.sql`
- `docker exec -it datareporter_postgres psql -h localhost -U app datareporter` pass:*app-my-super-pass*
- datareporter=>
- `\dt`
- `SELECT * FROM articles;`
- `\q`

#### SQLite
- `docker build . -t data-reporter-sqlite -f ./Docker/sqlite/Dockerfile`
- `docker run --rm data-reporter-sqlite sqlite3 --version`
- `docker run --rm -it -v "$(pwd)/db:/workspace" -w /workspace data-reporter-sqlite`
- sqlite>
- `.read ./db_script/sqlite.sql`
- `.databases`
- `.tables`
- `SELECT * FROM articles;`
- `.q`

### Warning
This is just an exercise, even though it works, it should not be used in production...
