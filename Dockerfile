FROM golang:1.10.2 as builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/sluongng/crud-example
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN go build -a -installsuffix -o /app .


FROM alpine:3.7
EXPOSE 7001
COPY --from=builder /app .
CMD ["./app"]
