#!/bin/sh

#由于无法在同一个docker-compose文件中获取其他容器的IP，并替换配置文件占位符
#所以这里将docker-compose文件拆开，在过程中通过shell脚本去执行该过程

chmod +x /bin/chatroom/replace-ip.sh

configName="config"$(date "+%Y%m%d%H%M%S")

/bin/chatroom/replace-ip.sh chatroom-rocketmq-namesrv /bin/chatroom/config.pro.yaml /bin/chatroom/$configName.yaml
/bin/chatroom/replace-ip.sh chatroom-rocketmq-broker /bin/chatroom/config.pro.yaml /bin/chatroom/$configName.yaml
/bin/chatroom/chatroom --config=/bin/chatroom/$configName.yaml

