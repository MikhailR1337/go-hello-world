FROM golang:alpine

WORKDIR /cmd

COPY main.go ./
COPY static/index.html ./static

RUN go build -o server main.go
EXPOSE 3000
ENTRYPOINT ["/cmd/server"]