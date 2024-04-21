FROM golang:latest

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go mod tidy && go build -o app ./cmd

CMD ["./app"]