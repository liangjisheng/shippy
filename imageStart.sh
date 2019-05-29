#!/bin/bash

cd $GOPATH/src/shippy/

# 如果想实时查看日志，可以使下面的每个 docker run 在一个终端运行,且去掉-d选项
docker run -d -p 5432:5432 postgres
docker run -d -p 27017:27017 mongo
docker run -d -p 4222:4222 nats

docker run -d --net="host" -e MICRO_REGISTRY=mdns user-service

docker run -d --net="host" -e MICRO_REGISTRY=mdns vessel-service

docker run -d --net="host" -e MICRO_REGISTRY=mdns consignment-service

# 最后启动,错误的token会认证不通过
# 这个token是通过user-cli代码中写的一个user调用user-service得到的
cd consignment-cli
docker run --net="host" -e MICRO_REGISTRY=mdns consignment-cli ./consignment-cli consignment.json eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiMjU0M2JmMjItZDBjOS00Y2I3LTllZjktZDdjZjU2MDQ4N2I3IiwibmFtZSI6ImxpYW5namlzaGVuZyIsImNvbXBhbnkiOiJsaWFuZ2ppc2hlbmciLCJlbWFpbCI6IjEyOTQ4NTE5OTBAcXEuY29tIiwicGFzc3dvcmQiOiIkMmEkMTAkMS9lckU2a0pwOHFYTThHWHdVZmx5dUM4eHVpOGguSjRRaUtwRWE1ZUdsRVE5SVVZRDVzdTIifSwiZXhwIjoxNTU5MjkwOTg1LCJpc3MiOiJnby5taWNyby5zcnYudXNlciJ9.y4ar2ZG8wePZ_rcuT_ejAzOZvMQqzCV1oDhTaQdWTMk
cd ..