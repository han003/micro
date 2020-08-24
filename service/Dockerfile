FROM golang:latest as builder

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download

ADD *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  

RUN apk --no-cache add curl

EXPOSE 9090

WORKDIR /app/

COPY --from=builder /app/app .
ADD migrations ./migrations

RUN ls

ENTRYPOINT ["/app/app"]
