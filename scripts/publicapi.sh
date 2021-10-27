#!/bin/bash

# Copy publicapi.json files from micro/services to m3o/m3o/api

pushd services

find . -name publicapi.json | while read file; do 
	name=`echo $file | sed 's@./\(.*\)/publicapi.json@\1@g'`;
	cp $file ../m3o/api/$name.json;
done

popd
