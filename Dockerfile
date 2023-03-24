FROM golang:1.20.2-alpine

WORKDIR /app

RUN go mod download

COPY . .

RUN go build -o /docker-authoraizer

expose 8080

CMD [ "/docker-authoraizer" ]
