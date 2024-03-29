#!/bin/bash

SERVER=$2
BASE_DIR=$PWD
INTERVAL=2

# 命令行参数，需要手动指定
ARGS=""

if [ "$SERVER" = "" ];then
	echo `"文件不存在" $SERVER`
  exit 1
fi

function start()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo "$SERVER already running"
		exit 1
	fi

	nohup $BASE_DIR/bin/$SERVER $ARGS  server &>/dev/null &

	echo "sleeping..." &&  sleep $INTERVAL

	# check status
	if [ "`pgrep $SERVER -u $UID`" == "" ];then
		echo "$SERVER start failed"
		exit 1
	fi
}

function status()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo $SERVER is running
	else
		echo $SERVER is not running
	fi
}

function stop()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		kill -9 `pgrep $SERVER -u $UID`
	fi

	echo "sleeping..." &&  sleep $INTERVAL

	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo "$SERVER stop failed"
		exit 1
	fi
}

function restart()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		kill -1 `pgrep $SERVER -u $UID`
	else
	  nohup $BASE_DIR/bin/$SERVER $ARGS  server &>/dev/null &
	fi

	echo "sleeping..." &&  sleep $INTERVAL

	# check status
	if [ "`pgrep $SERVER -u $UID`" == "" ];then
	  nohup $BASE_DIR/bin/$SERVER $ARGS  server &>/dev/null &
	  if [ "`pgrep $SERVER -u $UID`" == "" ];then
	    echo "$SERVER start failed"
	    exit 1
    fi
	fi
}

case "$1" in
	'start')
	start
	;;
	'stop')
	stop
	;;
	'status')
	status
	;;
	'restart')
	restart
	;;
	*)
	echo "usage: $0 {start|stop|restart|status}"
	exit 1
	;;
esac
