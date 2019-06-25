FROM golang:latest
MAINTAINER chendante
WORKDIR $GOPATH/src/Url-Shortener
ADD . $GOPATH/src/Url-Shortener/main.go
RUN go build .
EXPOSE 8080
ENTRYPOINT  ["./main"]