#!/bin/sh

reg_url=$1
version=$2

image=${reg_url}/httpserver:${version}

# 构建镜像
docker build -t $image .

# 提交镜像仓库
docker push $image

# 发布到k8s集群
sed -i -r "s#image:\s.*#image: ${image}#" httpserver-deploy.yaml
kubectl apply -f httpserver-deploy.yaml
