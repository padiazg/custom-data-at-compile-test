
https://goenning.net/2017/01/25/adding-custom-data-go-binaries-compile-time/


```
.GIT_COMMIT=$(shell git rev-parse HEAD)
.GIT_VERSION=$(shell git describe --tags --always --dirty 2>/dev/null)
.GIT_UNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
```


go build -ldflags "-X main.buildTime=$(date +'%Y.%m.%d.%H%M%S') -X main.commitHash=$(git rev-parse HEAD) -X main.version=$(git describe --tags --always --dirty 2>/dev/null)"