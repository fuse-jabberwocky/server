VERSION="latest"

docker tag jabberwocky-server:$VERSION  squakez/jbw:$VERSION
docker push squakez/jbw:$VERSION