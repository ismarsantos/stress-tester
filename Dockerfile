# Etapa de construção
FROM golang:1.18-alpine AS builder

WORKDIR /app

# Copiar go.mod e go.sum e baixar as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código-fonte para o diretório de trabalho
COPY . .

# Compilar o binário
RUN go build -o stress-tester .

# Verificar a presença do binário
RUN ls -l /app

# Etapa de execução
FROM alpine:latest

WORKDIR /root/

# Copiar o binário do estágio de construção
COPY --from=builder /app/stress-tester .

# Adicionar comando de listagem para verificar a presença do binário
RUN ls -l /root

CMD ["/root/stress-tester"]
