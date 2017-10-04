FROM golang:1.9

# Install deps
WORKDIR /go/src/github.com/lox/binarystar
ADD . /go/src/github.com/lox/binarystar
RUN go build -a .

EXPOSE 8924
ENTRYPOINT ["./binarystar"]
