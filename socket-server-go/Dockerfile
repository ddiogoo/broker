FROM golang:1.22.4

COPY . /go

RUN go get github.com/nats-io/nats
RUN go get github.com/joho/godotenv
RUN go get github.com/gorilla/websocket

RUN go build socket-server-go.go

EXPOSE 8080
ENTRYPOINT ["/go/socket-server-go"]
