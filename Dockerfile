FROM alpine:latest

RUN apk add --no-cache curl

WORKDIR /app
# Goは実質使っていないので、スクリプトのみコピー
COPY ./command/command/* .
