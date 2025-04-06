FROM golang:1.24.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /userManager ./cmd/userManager

FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=builder /userManager /userManager
COPY config /app/config
COPY database/migration /app/database/migration

EXPOSE 8080

ENTRYPOINT ["/userManager"]