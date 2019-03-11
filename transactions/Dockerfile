FROM golang

RUN mkdir -p /go/src/github.com/nknab/MonewayV1.0

ADD . /go/src/github.com/nknab/MonewayV1.0
WORKDIR /go/src/github.com/nknab/MonewayV1.0/transactions

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/nknab/MonewayV1.0/transactions/cmd  -watch github.com/nknab/MonewayV1.0/transactions
