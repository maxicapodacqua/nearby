FROM golang:1.20 as builder
WORKDIR /app
# Download Go modules
COPY go.mod go.sum ./
RUN go get ./...
COPY . ./
# Seeds database
RUN go run cmd/seed/sqlite.go
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /nearby-api cmd/nearby/main.go


FROM build AS run-test
# Run tests on docker build
RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release
WORKDIR /
COPY --from=builder /nearby-api /nearby-api
# Add database into image
COPY --from=builder /app/data/db.sqlite /data/db.sqlite
# Default port
EXPOSE 8080
USER nonroot:nonroot
CMD ["/nearby-api"]