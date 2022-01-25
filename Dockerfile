FROM golang
WORKDIR /usr/app/gobooks
COPY . .
CMD go run cmd/main.go
