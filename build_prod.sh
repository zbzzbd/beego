#!/bin/bash

source base.sh

cd $BASEDIR"/src"
go build -ldflags "-w" -o ${PROJECT_NAME}_${DATETIME}
