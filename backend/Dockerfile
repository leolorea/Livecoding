
FROM golang:1.21.1-alpine

RUN apk update && apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1

RUN go build -o main main .

EXPOSE 8080

CMD ["./main"]
