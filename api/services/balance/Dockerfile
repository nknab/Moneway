FROM golang

#RUN mkdir -p /go/src/github.com/nknab/Moneway

ADD . /go/src/github.com/nknab/Moneway
WORKDIR /go/src/github.com/nknab/Moneway/api/services/balance

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/nknab/Moneway/api/services/balance  -watch github.com/nknab/Moneway/api/services/balance