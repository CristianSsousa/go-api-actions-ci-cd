# Build stage
FROM golang:1.21-alpine AS builder

# Instalar dependências necessárias
RUN apk add --no-cache git

# Definir diretório de trabalho
WORKDIR /app

# Copiar go mod e sum files
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd/api

# Runtime stage
FROM alpine:latest

# Instalar ca-certificates para HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário do builder
COPY --from=builder /app/api .

# Expor porta (Cloud Run usa a variável PORT)
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./api"]

