FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/client -v cmd/client/client.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/client /client
COPY --from=builder /go/src/app/config /config
ARG target
RUN ["chmod", "+x", "/client"]
CMD ["/client", "-name", "${target}"]
