# reachability-tester

inspired by http://ec2-reachability.amazonaws.com

## quick go 

### web UI

```sh
docker run -d -v /path/to/config.json:/config/config.json -p 80:80 katie/reachability-ui
```

*See [example.json](https://github.com/kayteh/reachability-tester/blob/master/example.json) for an example config*

### nodes to test

```sh
docker run -d -p 80:80 katie/reachability-node
```

### docker sux?
if you don't like docker, try [github binaries](https://github.com/kayteh/reachability-tester/releases/latest)

web UI (app) binary takes an env var, `CONFIG_PATH` that points to the config.json.
both binaries (app, node) take an env var, `ADDR` that defines what IP:port (e.g. `127.0.0.1:8003`) it will be on.

## building/contributing

you need

- go (developed on 1.10)
- dep
- go-bindata: `go get -u github.com/jteeuwen/go-bindata/...`

```sh
dep ensure
go generate ./... # must be done every template or asset update
go build -o app ./cmd/app
go build -o node ./cmd/node
```

### abstract gist

node is a fasthttp server, only lives to serve a single file, but it's compiled in so it's a self-contained binary.

app is a fasthttp server, but renders an HTML template to serve as the test on once at start time. it also serves a single error image. the HTML template is generated from CONFIG_PATH's json.

web UI is a tachyon-based UI. nothing really more to it.