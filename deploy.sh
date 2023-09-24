#!/bin/bash

#由于无法在同一个docker-compose文件中获取其他容器的IP，并替换配置文件占位符
#所以这里将docker-compose文件拆开，在过程中通过shell脚本去执行该过程

docker-compose -f docker-compose-base.yaml  up --build -d

chmod +x ./replace-ip.sh && ./replace-ip.sh chatroom-rocketmq-namesrv ./config.pro.yaml ./config.yaml
chmod +x ./replace-ip.sh && ./replace-ip.sh chatroom-rocketmq-broker ./config.pro.yaml ./config.yaml

docker-compose -f docker-compose-myapp.yaml  up --build -d
