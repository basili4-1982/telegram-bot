FROM golang:1.25-alpine AS builder

WORKDIR /build

# Устанавливаем git и ca-certificates
RUN apk add --no-cache git ca-certificates

# Кешируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем код и собираем
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -trimpath -o tgbot .

# Финальный образ продолжается
FROM scratch

WORKDIR /app

# Копируем бинарник
COPY --from=builder /build/tgbot /app/tgbot

# Устанавливаем временную зону
ENV TZ=Europe/Moscow


CMD ["/app/tgbot"]