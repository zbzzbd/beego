#!/bin/bash

source base.sh

cd $BASEDIR"/src"
go build -o ${PROJECT_NAME}_${DATETIME}
