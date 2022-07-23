# Adding custom date at compile time

Reference
https://goenning.net/2017/01/25/adding-custom-data-go-binaries-compile-time/

build with
```bash
go build -ldflags "-X main.buildTime=$(date +'%Y.%m.%d.%H%M%S') -X main.commitHash=$(git rev-parse HEAD) -X main.version=$(git describe --tags --always --dirty)"
```