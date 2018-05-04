FROM golang:1.10

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

# build directories
RUN mkdir /app
RUN mkdir -p /go/src/ssoServer
ADD . /go/src/ssoServer
WORKDIR /go/src/ssoServer

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# Build my app
RUN go build -o /ssoServer/main .
CMD ["ssoServer/main"]