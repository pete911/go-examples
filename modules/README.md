# modules

## releasing

 - use [semver](https://semver.org)
 - check [module compatibility](https://blog.golang.org/module-compatibility)
 - run [gorelase](https://pkg.go.dev/golang.org/x/exp/cmd/gorelease)
 - check if relasing [major version](https://blog.golang.org/v2-go-modules)

## installing

Use go isntall instead of go get [module changes](https://blog.golang.org/go116-module-changes).
`go install <module>@<version>` e.g. `go isntall golang.org/x/tools/gopls@latest`