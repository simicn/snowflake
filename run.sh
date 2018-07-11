#!/bin/sh
echo "build and run snowflake"
export ETCD_HOST=http://192.168.3.15:2379
./build.sh
docker rm -f snowflake-1
docker run -it -d \
    --name snowflake-1 \
    -h snowflake-dev \
    -e MACHINE_ID=1 \
    -e ETCD_HOSTS=$ETCD_HOST \
    -p 10000:10000 \
    snowflake
