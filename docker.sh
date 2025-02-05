#!/bin/bash
#!/bin/bash

# 检查当前用户是否为 root
if [ "$(id -u)" != "0" ]; then
   echo "此脚本必须以 root 用户权限运行，请使用 sudo 执行此脚本。" 1>&2
   exit 1
fi

# 获取所有监听3306端口的进程ID
pids=$(netstat -tulnp | grep :3306 | awk '{print $7}' | sed 's/\/.*$//')

# 循环遍历所有PID，并杀掉它们
for pid in $pids; do
    echo "Killing process with PID: $pid"
    kill $pid 2>/dev/null
done

# 等待一段时间，确保进程被杀掉
sleep 2

# 启动Docker容器
docker compose up -d
