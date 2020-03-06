# Flag

This example has two flags, `host` and `port`. Both flags can be set using env variables and have default values as well.
There is test example as well that can be run with `go test -v ./...`.

 - default values `go run .`
 - set flags `go run . --host test --port 443`
 - use env. vars `GOX_HOST=staging GOX_PORT=8443 go run .`
 - override env. var with flag `GOX_HOST=staging GOX_PORT=8443 go run . --port 9000`
