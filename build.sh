#!/bin/sh
echo "build snowflake"
docker build --no-cache --rm=true -f Dockerfile.dev -t snowflake .