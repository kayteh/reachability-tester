FROM golang:alpine
WORKDIR /go/src/github.com/kayteh/reachability-tester 
COPY . /go/src/github.com/kayteh/reachability-tester
RUN go build ./cmd/app

FROM alpine
COPY --from=0 go/src/github.com/kayteh/reachability-tester/app /app
COPY example.json /config/config.json
ENV CONFIG_PATH=/config/config.json
CMD /app