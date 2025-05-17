FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./src/

RUN go build -o oolio-api-ecommerce ./src/cmd/main.go

EXPOSE 8080

CMD ["./oolio-api-ecommerce"]