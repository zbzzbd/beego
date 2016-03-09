#!/bin/bash
BASEDIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
DATETIME=`date +%Y%m%d%H%M%S`
PROJECT_NAME="achilles"

if [[ $GOPATH == *$BASEDIR* ]]
then
	echo $GOPATH
else
	export GOPATH=$GOPATH":"$BASEDIR
	echo $GOPATH
fi
