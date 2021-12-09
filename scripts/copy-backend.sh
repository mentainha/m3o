#!/bin/bash

set -e
set -x

DIR=$(realpath `dirname $0`)
PWD=`pwd`

git clone https://github.com/m3o/backend /tmp/backend

cd /tmp/backend

grep -r "github.com/micro/services" * | cut -f 1 -d : | sort | uniq | \
xargs -r sed -i 's@github.com/micro/services@github.com/m3o/m3o/services@g'

grep -r "github.com/m3o/services" * | cut -f 1 -d : | sort | uniq | \
xargs -r sed -i 's@github.com/m3o/services@github.com/m3o/m3o/backend@g'

grep -r "github.com/m3o/backend" * | cut -f 1 -d : | sort | uniq | \
xargs -r sed -i 's@github.com/m3o/backend@github.com/m3o/m3o/backend@g'

rm go.mod
go mod init github.com/m3o/m3o/backend
go mod tidy

rsync -avz --delete --exclude=README.md --exclude=.git --exclude=.github --exclude=CNAME /tmp/backend/ $DIR/../backend/

cd $PWD

rm -rf /tmp/backend
