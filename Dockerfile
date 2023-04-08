FROM golang:1.20.2-alpine

WORKDIR /app

RUN go mod download

COPY pkg .

COPY main.go .

RUN go build -o /docker-eligibility

expose 8080

CMD [ "/docker-eligibility" ]
