FROM golang:1.20-alpine3.18

ADD ./srcs /ft_auth_bot/
WORKDIR /ft_auth_bot

RUN go install .

CMD ft_auth_bot
