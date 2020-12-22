VERSION="latest"

if [ -z "$1" ] ; then
    echo "No argument supplied, building the docker image as LATEST"
else
    VERSION="$1"
    echo "Building the docker image as $VERSION"
fi

docker tag jbw-server:$VERSION  squakez/jbw-server:$VERSION
docker push squakez/jbw-server:$VERSION