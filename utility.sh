#!/bin/bash

COMMAND="$1"

function deps() {
         rm -rf vendor
         go get ./...
         govendor init
         govendor add +external
}

 if [ "$COMMAND" = "run" ];
        then
                echo "run"
                govendor fetch +outside
                docker build -t nafue-api .
                docker run --publish 9090:8080  --name nafue-api --env-file .env --rm nafue-api

fi
if [ "$COMMAND" = "build" ]; then
            govendor fetch +outside
            rm -rf dist/
            echo "build"
            mkdir dist
            zip -r dist/build.zip ./ -x \*node_modules/* \*www/* \*.git/* \.env \\\.idea/* *\bower_components/* \*dist/*
fi

if [ "$COMMAND" = "deps" ]; then
           echo "manage deps"
           deps
fi


