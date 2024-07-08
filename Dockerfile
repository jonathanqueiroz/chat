# Imagem base
FROM golang:1.22.3-alpine

# Diretório de trabalho
WORKDIR /app

# Copia o código da aplicação para a imagem
COPY . .

# Instala as dependências
RUN go mod download

# Compila a aplicação
RUN go build -o main ./cmd/server/main.go

# Expõe a porta 8080
EXPOSE 8080

# Comando para executar a aplicação
CMD ["/app/main"]
