VERSION="latest"

CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o build/_output/bin/jabberwocky ./cmd/server/main.go

docker build -t jabberwocky-server:$VERSION -f build/Dockerfile .