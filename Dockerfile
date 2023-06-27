FROM golang:1.20-alpine3.18

ADD bot /diablo_iv_tool/
WORKDIR /diablo_iv_tool

RUN apk add gcc g++

RUN mv ./cmd ./dbiv_tool
RUN CGO_ENABLED=1 GOOS=linux go install ./dbiv_tool

CMD dbiv_tool
