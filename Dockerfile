FROM golang:1.4.2

# Install deps
WORKDIR /go/src/github.com/lox/binarystar
RUN go get github.com/tools/godep
ADD Godeps /go/src/github.com/lox/binarystar/Godeps
RUN godep restore

# Add source and build
ADD . /go/src/github.com/lox/binarystar
RUN go build -a .

EXPOSE 8924
ENTRYPOINT ["./binarystar"]