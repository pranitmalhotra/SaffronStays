FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o saffronstays-api .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/saffronstays-api .
ENV PORT=8000
EXPOSE 8000
CMD ["./saffronstays-api"]
