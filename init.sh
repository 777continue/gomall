#!/bin/bash

# 定义数据库初始化脚本的路径
DB_INIT_SCRIPT="./init_db.sql"

# 检查脚本文件是否存在
if [ ! -f "$DB_INIT_SCRIPT" ]; then
    echo "Database initialization script not found."
    exit 1
fi
# 连接到 MySQL 数据库
mysql -h 127.0.0.1 -P 3306 -u root -p'root' -e "source $DB_INIT_SCRIPT"

echo "Database init completed."