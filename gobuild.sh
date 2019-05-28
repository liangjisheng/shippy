#!/bin/bash

cd consignment-cli
go build
cd ..

cd consignment-service
go build
cd ..

cd vessel-service
go build
cd ..

cd user-cli
go build
cd ..

cd user-service
go build
cd ..
