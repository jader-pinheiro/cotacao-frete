# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copia go.mod e go.sum para cache de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código
COPY . .

# Instala o Swag CLI e gera a documentação
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

# Compila a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copia binário compilado
COPY --from=builder /app/server .+

# Copia a pasta docs gerada pelo swag
COPY --from=builder /app/docs ./docs

# Instala tzdata para timezone
RUN apk add --no-cache tzdata

EXPOSE 3000

# Inicia a aplicação
CMD ["./server", "api", "run"]
