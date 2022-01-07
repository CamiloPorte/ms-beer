# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /MsBeer
COPY . .

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN go build -o /docker-ms-beer cmd/main.go
##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build  /docker-ms-beer  /docker-ms-beer

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-ms-beer"]