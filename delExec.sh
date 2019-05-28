#!/bin/bash

if [ -f ./consignment-cli/consignment-cli ]
then
    rm ./consignment-cli/consignment-cli
fi

if [ -f ./consignment-service/consignment-service ]
then
    rm ./consignment-service/consignment-service
fi

if [ -f ./vessel-service/vessel-service ]
then
    rm ./vessel-service/vessel-service
fi

if [ -f ./user-cli/user-cli ]
then
    rm ./user-cli/user-cli
fi

if [ -f ./user-service/user-service ]
then
    rm ./user-service/user-service
fi