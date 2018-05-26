FROM golang:1.10.2 as builder

# Download and install the latest release of Go Dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/sluongng/crud-example
# Run Go Dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
# Run Go Build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

# Deploy
FROM alpine:3.7
COPY --from=builder /app .
ENTRYPOINT ["./app"]
EXPOSE 7001
