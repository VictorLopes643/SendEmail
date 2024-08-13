# Usa uma imagem base do Golang
FROM golang:1.22-alpine

# Define o diretório de trabalho
WORKDIR /app

# Copia o código-fonte para o contêiner
COPY . .

# Baixa as dependências
RUN go mod tidy

# Compila a aplicação
RUN go build -o main ./internal/main.go

# Exponha a porta que a aplicação vai usar
EXPOSE 8080

# Define o comando padrão para iniciar o binário
# CMD ["./main"]
CMD ["tail", "-f", "/dev/null"]

