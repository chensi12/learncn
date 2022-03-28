# 模块十一作业
安装了istio
配置了httpserver 服务以 Istio Ingress Gateway 的形式发布
创建了tls secret，以https访问服务
![image](https://user-images.githubusercontent.com/19473836/160321150-7a7c3917-f9e3-416f-846f-cf4c9d9e9cb4.png)
![image](https://user-images.githubusercontent.com/19473836/160321195-04eca6ef-0f07-460a-9eb4-f7e609fbcc28.png)
![image](https://user-images.githubusercontent.com/19473836/160363329-6471fc91-01c4-4702-97a4-cc67082bfeb2.png)


# 模块十作业
为 HTTPServer 添加 0-2 秒的随机延时；
![image](https://user-images.githubusercontent.com/19473836/158147889-95632269-59fc-426c-95e4-fb59760fb79a.png)

为 HTTPServer 项目添加延时 Metric；
![image](https://user-images.githubusercontent.com/19473836/158147991-16653999-f342-47a0-9b0f-b8a0a86ec332.png)

将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；
httpserver-deploy.yaml

从 Promethus 界面中查询延时指标数据；
![image](https://user-images.githubusercontent.com/19473836/158151911-96e46aeb-247b-47f8-ba1f-5c462f2f9b28.png)

（可选）创建一个 Grafana Dashboard 展现延时分配情况。
