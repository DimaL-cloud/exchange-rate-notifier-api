FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/main.go
EXPOSE 8080
ENTRYPOINT ["/app/main"]