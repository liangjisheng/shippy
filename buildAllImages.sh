#!/bin/bash

cd consignment-cli
make build
cd ..

cd consignment-service
make build
cd ..

cd vessel-service
make build
cd ..

cd user-cli
make build
cd ..

cd user-service
make build
cd ..

cd email-service
make build
cd ..
