VERSION="latest"

if [ -z "$1" ] ; then
    echo "No argument supplied, building the docker image as LATEST"
else
    VERSION="$1"
    echo "Building the docker image as $VERSION"
fi

CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o build/_output/bin/jabberwocky ./cmd/server/main.go

docker build -t jbw-server:$VERSION -f build/Dockerfile .