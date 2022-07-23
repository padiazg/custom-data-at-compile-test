# Adding custom date at compile time

Reference
https://goenning.net/2017/01/25/adding-custom-data-go-binaries-compile-time/

# Description
The ides is to set the values for `buildTime`, `commitHash` and `version` at compile time using the **-ldflags** compiler flag

## source of values
### version
Taken from the last tag set
```bash
git tag v0.0.1
git describe --tags --always --dirty
echo $(git describe --tags --always --dirty)
```

### commitHash
Here we look for the last commit hash. There are several ways to get it
```bash
# shot version
$ git log -1 --format='%h'
149179c

# long version
$ git log -1 --format='%H'
149179c16ae7e57129a1be076101d04ac34c6222

# long version, another approach
$ git rev-parse HEAD
149179c16ae7e57129a1be076101d04ac34c6222
```

## Build
```bash
go build -ldflags "-X main.buildTime=$(date +'%Y.%m.%d.%H%M%S') -X main.commitHash=$(git rev-parse HEAD) -X main.version=$(git describe --tags --always --dirty)"
```