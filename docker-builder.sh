#!/usr/bin/env bash

REPO=${1}
ENVIRONMENT=${2}

#Set Defaults
if [ "$REPO" = "" ]; then
    REPO="local"
fi

if [ "$ENVIRONMENT" = "" ]; then
    ENVIRONMENT="dev"
fi

# Set Options for Docker Build
if [ "$ENVIRONMENT" = "hash" ]; then
    docker build -t ${REPO}/primecalc:$(git rev-parse --short=8 HEAD) ./
    if [ $? -eq 0 ]; then
        echo "Success"
    else
        echo "Failed!"
    fi
elif [ "$ENVIRONMENT" = "dev" ]; then
    docker build -t ${REPO}/primecalc:dev-build ./
    if [ $? -eq 0 ]; then
        echo "Success"
    else
        echo "Failed"
    fi
elif [ "$ENVIRONMENT" = "prod" ]; then
    docker build -t ${REPO}/primecalc:latest ./
    if [ $? -eq 0 ]; then
        echo "Success"
    else
        echo "Failed"
    fi
else
    docker build -t ${REPO}/primecalc:${ENVIRONMENT} ./
    if [ $? -eq 0 ]; then
        echo "Success"
    else
        echo "Failed"
    fi
fi
