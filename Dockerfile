FROM golang:1.21.3-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o bin .

FROM alpine:3
ENV APP_PORT=3000
WORKDIR /app
COPY --from=builder bin .
ENTRYPOINT [ "bin" ]

EXPOSE ${APP_PORT}