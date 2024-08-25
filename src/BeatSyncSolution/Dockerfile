FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN CGO_ENABLED=0 go build -o /beatsyncsolution

FROM python:3.9 AS base

RUN apt-get update && apt-get -y install ffmpeg git

RUN python3 -m pip install --user pipx
RUN python3 -m pipx ensurepath
RUN python3 -m pipx install https://get.zotify.xyz

WORKDIR /app

COPY --from=builder /beatsyncsolution ./beatsyncsolution

EXPOSE 3000

CMD [ "./beatsyncsolution" ]


