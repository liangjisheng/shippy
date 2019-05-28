#!/bin/bash

CONTAINTERS=$(docker ps -aq)
if [ "$CONTAINTERS" != "" ]
then
    docker rm -f $CONTAINTERS
fi