FROM golang:latest as builder

WORKDIR /go/src/app

RUN go get -d -v github.com/gorilla/mux \
	&& go get -d -v gopkg.in/mgo.v2/bson \
	&& go get -d -v gopkg.in/mgo.v2

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  

RUN apk --no-cache add curl

EXPOSE 9090

WORKDIR /root/

COPY --from=builder /go/src/app .

CMD ["./app"]  