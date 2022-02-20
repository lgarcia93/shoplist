# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS build

RUN apk add -U --no-cache ca-certificates

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./ ./

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go mod vendor

RUN swag init -g cmd/shoplist/main.go  --parseDependency

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /docker-go-shoplist ./cmd/shoplist/main.go

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /

COPY --from=build /docker-go-shoplist /docker-go-shoplist

EXPOSE 5000

ENV shoplist_api_env=prod

ENTRYPOINT ["/docker-go-shoplist"]
