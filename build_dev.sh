#!/bin/bash

source base.sh

cd $BASEDIR"/src"
go build -gcflags "-N -l" -o $PROJECT_NAME

