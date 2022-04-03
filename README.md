# some-protos-go-example

Sample app that spins up a gRPC server with protobufs built from
[kevinmichaelchen/some-protos-go](https://github.com/kevinmichaelchen/some-protos-go).

This server should run in conjunction with the NodeJS TypeScript client at
[some-protos-ts-example](https://github.com/kevinmichaelchen/some-protos-ts-example).

## Running
Install dependencies with `go mod download` and then run
```bash
make
```