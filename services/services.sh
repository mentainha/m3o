#!/bin/bash

# This script builds a top level client for services

DIR=`pwd`

rm -rf /tmp/services
git clone https://github.com/micro/services /tmp/services
cd /tmp/services

SERVICES=`find . -maxdepth 2 -type d -name proto | cut -f 2 -d / | sort`

## rename everything
grep -r github.com/micro/services | cut -f 1 -d : | sort | uniq | xargs sed -i 's@github.com/micro/services@m3o.dev/services@g'

## copy everything first
for service in ${SERVICES[@]}; do
	if [ ! -d $DIR/$service ]; then
		mkdir $DIR/$service
	fi

	rsync -q -avz --delete $service/ $DIR/$service/
done

## copy any top level go files
cp *.go $DIR/

cd $DIR

go fmt ./...
