#! /bin/bash

MODULE_NAME="agent"
DIR="$( cd "$( dirname "$0"  )" && pwd  )"
MODULE_PATH=$DIR/$MODULE_NAME

ps -ef|grep $MODULE_PATH |grep -v grep | awk -F' ' '{print $8}'
ps -ef|grep $MODULE_PATH |grep -v grep | awk -F' ' '{print $2}'|xargs kill -9
