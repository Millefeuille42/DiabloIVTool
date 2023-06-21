FROM golang:1.20-alpine3.18

ADD ./srcs /diablo_iv_tool/
WORKDIR /diablo_iv_tool

RUN go install .

CMD diablo_iv_tool
