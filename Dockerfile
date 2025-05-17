FROM golang:1.24.2 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh

COPY . .
RUN go build -o cmd/app/main

FROM ubuntu:22.04

WORKDIR /app
RUN apt-get update && apt-get install -y bash

COPY --from=builder /app/cmd/app/main .
COPY --from=builder /app/wait-for-it.sh ./wait-for-it.sh

COPY dbconfig/config.json /app/dbconfig/config.json
RUN chmod +x ./main ./wait-for-it.sh

EXPOSE 8080
CMD ["./wait-for-it.sh", "db:3306", "--", "./main"]