# API 网关现在运行在 8080 端口，同时告诉它和其他微服务一样使用 mdns 做服务发现
# 最后使用的命名空间是 shippy，它会作为我们服务名的前缀，比如 shippy.auth，
# shippy.email，默认值是 go.micro.api，如果不指定而使用默认值将无法生效
run:
	docker run -p 8080:8080 \
    	-e MICRO_REGISTRY=mdns \
    	microhq/micro api \
    	--handler=rpc \
    	--address=:8080 \
    	--namespace=shippy
