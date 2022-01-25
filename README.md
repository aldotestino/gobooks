# Gobooks

Simple API made with Go

## How to run

```sh
go run cmd/main.go
```

## How to build

```sh
go build cmd/main.go
```

## How to run with Docker

```sh
docker compose up
```

## Add a book

```sh
curl -X POST -H "Content-Type: application/json" \ 
  -d '{"Title": "1984", "Author": "George Orwell"}' \ 
  http://localhost:8080/books
```

## Get a book

```sh
curl  http://localhost:8080/book/[id]
```
## Get all books

```sh
curl  http://localhost:8080/books
```

