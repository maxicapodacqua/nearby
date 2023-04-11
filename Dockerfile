FROM golang:1.20

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go get ./...

COPY . ./
# Run tests on docker build
RUN go test ./...
# Seeds database
RUN go run cmd/seed/sqlite.go
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /nearby-api cmd/nearby/main.go

# Default port
EXPOSE 8080

CMD ["/nearby-api"]