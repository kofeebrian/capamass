FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN go get -v github.com/OWASP/Amass/v3/...
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v cmd/service/service.go

FROM alpine AS production
WORKDIR /
COPY --from=builder /go/bin/amass /bin
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/utils/config /utils/config
EXPOSE 3000
ENTRYPOINT [ "/app" ]
