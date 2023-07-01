FROM golang:1.20-alpine3.18

ADD bot /bot
WORKDIR /bot

RUN apk add gcc g++

RUN CGO_ENABLED=1 GOOS=linux go install .

CMD bot
