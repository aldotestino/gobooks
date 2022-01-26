FROM golang
WORKDIR /usr/app/gobooks
COPY . .
CMD go run src/cmd/main.go
