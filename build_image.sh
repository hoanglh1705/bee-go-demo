#!/bin/sh
TAG=$(date '+%Y%d%m.%H%M%S')
ImageName="lhhoangit/bee-go-demo"

echo "Building image $ImageName:$TAG"
docker build . -t $ImageName:$TAG
docker tag $ImageName:$TAG $ImageName:latest
echo "Push image $ImageName:$TAG"
docker push $ImageName:$TAG
docker push $ImageName:latest