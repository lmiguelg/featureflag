FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./

RUN go mod download

COPY . .

RUN apk update

RUN apk add make

RUN make build

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]