FROM golang:alpine
WORKDIR /go/src/github.com/kayteh/reachability-tester 
COPY . /go/src/github.com/kayteh/reachability-tester
RUN go build ./cmd/node

FROM alpine
COPY --from=0 go/src/github.com/kayteh/reachability-tester/node /node
CMD /node