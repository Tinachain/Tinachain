#! /bin/bash

MODULE_NAME="agent"
DIR="$( cd "$( dirname "$0"  )" && pwd  )"
MODULE_PATH=$DIR/$MODULE_NAME

ps -ef | grep $MODULE_PATH | grep -v grep
if [ $? -ne 0 ]
then
echo "start $MODULE_PATH....."
setsid $MODULE_PATH &
else
echo "$MODULE_PATH runing....."
fi