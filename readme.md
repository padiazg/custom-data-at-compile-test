
https://goenning.net/2017/01/25/adding-custom-data-go-binaries-compile-time/

go build -ldflags “-X main.buildTime=$(date +”%Y.%m.%d.%H%M%S”) -X main.commitHash=$(git log -1 --format='%h')”

