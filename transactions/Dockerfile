FROM golang

RUN mkdir -p /go/src/github.com/nknab/Moneway

ADD . /go/src/github.com/nknab/Moneway
WORKDIR /go/src/github.com/nknab/Moneway/transactions

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/nknab/Moneway/transactions/cmd  -watch github.com/nknab/Moneway/transactions
