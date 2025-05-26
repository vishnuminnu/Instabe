FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -o myapp

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/myapp .
CMD ["./myapp"]
