To install new go module

```shell
go mod tidy
```

To re-direct the Go tools for searching public module to local module

```shell
go mod edit -replace [expected-import-statement]=[dir-path-to-local-module]

# please checkout 02_go_modules/hello for example
go mod edit -replace example.com/greetings=../greetings
```

To run unit test with detail

```shell
go test -v #-v --> verbose
```

Get go documentation

```shell
godoc -http=localhost:8000
```

Get alert of any subtle bug

```shell
go vet
```
