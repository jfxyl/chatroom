#!/bin/bash

# 检查是否传递了容器名称作为参数
if [ -z "$1" ]; then
    echo "Usage: $0 <container_name>"
    exit 1
fi
# 原配置文件
if [ -z "$2" ]; then
    echo "Usage: $1 <config_old_path>"
    exit 1
fi
# 如果存在，则将原配置文件copy一份到新文件，并修改新文件，保证原配置文件始终有占位符，应对多次部署容器ip发生变化的情况
if [ -z "$3" ]; then
    echo "Usage: $2 <config_new_path>"
    exit 1
fi

# 获取传递的容器名称
CONTAINER_NAME="$1"
# 原配置文件
CONFIG_OLD_PATH="$2"
# 新配置文件
CONFIG_NEW_PATH="$3"
# 获取容器的 IP 地址
IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' "$CONTAINER_NAME")

#copy一个新配置文件
cp $CONFIG_OLD_PATH $CONFIG_NEW_PATH

echo $CONTAINER_NAME
echo $CONFIG_OLD_PATH
echo $CONFIG_NEW_PATH
echo $IP

# 替换文件中的字符串
sed -i "s/{$CONTAINER_NAME}/$IP/g" $CONFIG_NEW_PATH
