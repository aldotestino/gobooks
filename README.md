# Gobooks

Simple API made with Go

## How to run

1. Build (*)
```sh
docker compose build
```

2. Start 
```sh
docker compose up
```

3. Enter `db` container and create `gobooksdb` database (*)
```sh
docker exec -it db psql -U postgres
```

```sql
CREATE DATABASE gobooksdb;
```

4. Enter `gobooks-api` container and run the migration (*)
```sh
docker exec -it gobooks-api /bin/bash
```

```sh
go run src/cmd/db/migrate.go
```

(*) Run the steps  [1, 3, 4] only the first time

## Usage

### Add a book

```sh
curl -X POST -H "Content-Type: application/json" \ 
  -d '{"Title": "1984", "Author": "George Orwell"}' \ 
  http://localhost:8080/books
```

### Get a book

```sh
curl  http://localhost:8080/book/[id]
```
### Get all books

```sh
curl  http://localhost:8080/books
```

