# syntax=docker/dockerfile:1

FROM golang:1.17 as build

WORKDIR /app
COPY ./ ./

RUN go mod download
RUN go build ./cmd/apiq/main.go

FROM ubuntu

WORKDIR /

EXPOSE 8080

COPY --from=build /app/main /main
COPY ./config.yaml /config.yaml
COPY ./config.local.yaml /config.local.yaml
CMD ["/main"]