FROM golang:latest
MAINTAINER PATihomirov
LABEL version="1.0"
RUN mkdir "/app"
ADD . /app
WORKDIR /app
RUN go build -o main ./cmd
CMD ["/app/main"]