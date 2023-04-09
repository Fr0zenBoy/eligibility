FROM golang:1.20.2-alpine

WORKDIR /usr/src/app

COPY . .

RUN go mod download

RUN go build -o /docker-eligibility

expose 8080

CMD [ "/docker-eligibility" ]
