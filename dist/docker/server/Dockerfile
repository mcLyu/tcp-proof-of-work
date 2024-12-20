FROM golang:1.23.0 AS builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main_server ./cmd/server/main_server.go
RUN go test -v ./...

FROM scratch
COPY --from=builder /app /app
CMD ["/app/main_server"]