FROM golang:1.22.2-alpine3.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -o app ./cmd/app/main.go

EXPOSE 8080

CMD [ "./app" ]