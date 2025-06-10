# Etapa de build
FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Compila binário totalmente estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o servidor

# Etapa final - imagem mínima
FROM alpine:3.18
COPY --from=builder /app/servidor /servidor
ENTRYPOINT ["/servidor"]