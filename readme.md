# Adding custom date at compile time

Reference
https://goenning.net/2017/01/25/adding-custom-data-go-binaries-compile-time/

# Description
The idea is to set the values for `buildTime`, `commitHash` and `version` at compile time using the **-ldflags** compiler flag

## Source of values
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

## Build and run
```bash
go build -ldflags "-X main.buildTime=$(date +'%Y.%m.%d.%H%M%S') -X main.commitHash=$(git rev-parse HEAD) -X main.version=$(git describe --tags --always --dirty)"
```
Run
```bash
# the version is "dirty" as there were updates not commited
$ ./custom-data-at-compile
Build Time: 2022.07.22.213219
Commit Hash: 3fc0c356295f442c4f272c78604c0f5669bedcc9
Version: v0.0.1-3-g3fc0c35-dirty
Formatted: v0.0.1-3-g3fc0c35-dirty       SHA: 3fc0c356295f442c4f272c78604c0f5669bedcc9   Build date: 2022.07.22.213219 
# clean repo run, no dirty version
$ ./custom-data-at-compile
Build Time: 2022.07.22.215148
Commit Hash: 0361d4240a60a27a52f01c3defda2e8b1b5b7d20
Version: v0.0.1-7-g0361d42
Formatted: v0.0.1-7-g0361d42     SHA: 0361d4240a60a27a52f01c3defda2e8b1b5b7d20   Build date: 2022.07.22.215148  
```
