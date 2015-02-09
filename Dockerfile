FROM golang:1.3-cross
RUN apt-get update && apt-get install -y --no-install-recommends openssh-client
RUN go get github.com/mitchellh/gox
RUN go get github.com/aktau/github-release
RUN go get github.com/tools/godep

# ENV GOPATH /go/src/github.com/nlamirault/go-onlinelabs/Godeps/_workspace:/go/
ENV GOPATH /go/
WORKDIR /go/src/github.com/nlamirault/go-onlinelabs

ADD src/github.com/nlamirault/go-onlinelabs /go/src/github.com/nlamirault/go-onlinelabs
ADD Godeps/_workspace/ /go/
