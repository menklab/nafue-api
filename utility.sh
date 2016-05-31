#!/bin/bash

COMMAND="$1"

function deps() {
         rm -rf Godeps vendor
         go get -u github.com/aws/aws-sdk-go
         go get ./...
         export GO15VENDOREXPERIMENT=1
         godep save ./...
         export GO15VENDOREXPERIMENT=0
}

 if [ "$COMMAND" = "run" ];
        then
                echo "run"
                docker build -t moja-portal .
                docker run --publish 9090:8080  --name moja-portal --env-file .env --rm moja-portal

fi
if [ "$COMMAND" = "build" ]; then
            echo "manage deps"
            rm -rf dist/
            deps
            echo "build"
            mkdir dist
            zip -r dist/build.zip api config database models repositories utility services .dockerignore Dockerfile Dockerrun.aws.json main.go utility.sh
fi

if [ "$COMMAND" = "deps" ]; then
           echo "manage deps"
           deps
fi


