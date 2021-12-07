#!/bin/bash

set -e
set -x

DIR=$(realpath `dirname $0`)
PWD=`pwd`

git clone https://github.com/micro/services /tmp/services

cd /tmp/services

grep -r "github.com/micro/services" * | cut -f 1 -d : | sort | uniq | \
xargs sed -i 's@github.com/micro/services@github.com/m3o/m3o/services@g'

go mod tidy

rsync -avz --delete --exclude=README.md --exclude=.git --exclude=.github --exclude=CNAME /tmp/services/ $DIR/../services/

cd $PWD

rm -rf /tmp/services 
