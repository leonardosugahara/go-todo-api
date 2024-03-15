# Usar a imagem oficial do Go como base
FROM golang:1.22.1-alpine AS builder

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o código fonte da aplicação para o diretório de trabalho
COPY . .

# Compilar a aplicação
RUN go build -o todo_api main.go

# Segunda etapa para criar uma imagem menor
FROM debian:buster-slim

# Copiar o binário compilado da primeira etapa
COPY --from=builder /app/todo_api /usr/local/bin/todo_api

# Expor a porta 8080 para acessar a aplicação
EXPOSE 8080

# Comando para rodar a aplicação quando o contêiner for iniciado
CMD ["todo_api"]