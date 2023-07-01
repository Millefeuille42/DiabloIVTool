FROM golang:1.20-alpine3.18

ADD fetcher /fetcher
WORKDIR /fetcher

RUN go install .

CMD fetcher
