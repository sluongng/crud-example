FROM golang:1.10.2-alpine3.7 as builder

RUN apk add --no-cache curl git
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/sluongng/crud-example
ADD . /go/src/github.com/sluongng/crud-example

WORKDIR /go/src/github.com/sluongng/crud-example
RUN dep ensure -vendor-only
RUN go build -o /app main.go


FROM alpine:3.7
EXPOSE 7001
CMD ["./app"]
COPY --from=builder /app .
