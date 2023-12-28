# Use a imagem oficial do Golang como base
FROM golang:1.18

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para baixar as dependências
COPY go.mod go.sum ./

# Execute o comando go mod download para baixar as dependências
RUN go mod download

# Copie o código-fonte do aplicativo para o diretório de trabalho no contêiner
COPY . .

# Compile o aplicativo Go
RUN go build -o main .

# Exponha a porta em que o aplicativo está ouvindo (ajuste conforme necessário)
EXPOSE 5000
EXPOSE 5001

# Defina o comando padrão para iniciar o aplicativo
CMD ["./main"]
