#!/bin/bash

source base.sh

pgrep $PROJECT_NAME | xargs kill -9

cd $BASEDIR"/src"
./$PROJECT_NAME --env=test



