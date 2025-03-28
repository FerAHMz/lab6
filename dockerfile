FROM golang:1.21-alpine

WORKDIR /app

RUN apk add --no-cache gcc musl-dev sqlite-dev

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8081

CMD ["go", "run", "main.go"]