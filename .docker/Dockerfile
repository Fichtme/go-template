FROM golang:1.13.6-alpine as base
LABEL maintainer="Henny krijnen <krijnen.h@gmail.com>"
ARG APP_ENV
WORKDIR /app

FROM base as watcher
RUN apk add --no-cache\
    git \
    bash \
    vim
RUN go get github.com/codeskyblue/fswatch
COPY .docker/.fsw.yml .fsw.yml
RUN cat .fsw.yml
COPY main /go/src/main
CMD ["fswatch"]

FROM base as build
COPY main/go.mod main/go.sum ./
RUN go mod download
COPY main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest as release
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=base /app/main .
EXPOSE 8080
CMD ["./main"]


