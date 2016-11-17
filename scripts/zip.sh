#!/bin/bash

WORKING_DIR=$1

pushd . 2>&1 /dev/null
cd $WORKING_DIR
find . -type d -print -exec zip -r '{}'.zip ./'{}' \;
ls -1 | grep -v zip | xargs rm -rf
popd . 2>&1 /dev/null
