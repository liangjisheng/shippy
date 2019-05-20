# vessel micro service

consignment-service 负责记录货物的托运信息，现在创建第二个微服务 vessel-service 来选择合适的货轮来运送货物

consignment.json 文件中的三个集装箱组成的货物，目前可以通过 consignment-service 管理货物的信息，现在用 vessel-service 去检查货轮是否能装得下这批货物

现在需要修改 consignent-service/main.go，使其作为客户端去调用 vessel-service，查看有没有合适的轮船来运输这批货物

至此，我们完整的将 consignment-cli，consignment-service，vessel-service 三者流程跑通了

客户端用户请求托运货物，货运服务向货船服务检查容量、重量是否超标，再运送