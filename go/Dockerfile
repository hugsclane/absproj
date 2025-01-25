FROM golang:1.23.5 AS build

WORKDIR /app

COPY --link go.mod go.sum ./

RUN go mod download

COPY --link . .

RUN go build -o main ./cmd

FROM alpine:3.21.2

WORKDIR /app

COPY --link --from=build /app/main ./main

ENTRYPOINT [ "./main" ]