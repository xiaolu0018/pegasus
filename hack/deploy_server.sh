#!/usr/bin/env bash
runDirName=pegasus_running

function init_Dir() {
   if [ -d ~/${runDirName} ]; then
      rm -rf ~/${runDirName}
   fi

    mkdir ~/${runDirName}
}

function clean() {
    cd ~
    rm pegasus
    rm dist.tar.gz
}

function target::Kill_old() {
    pid=$(pgrep -u weixin pegasus_$1)
    kill -9 $pid
}

function target::Prepare() {
    cp pegasus pegasus_$1
    mv pegasus_$1 ~/${runDirName}/
}

function target::Start() {
    cd ~/${runDirName}
    ./pegasus_$1 $1 start> ./$1.log 2>&1 &
}

function startTarget() {
    cd ~
    target::Kill_old $1
    target::Prepare $1
    target::Start $1
}

function deployDist() {

    if [ -d ~/${runDirName}/dist ]; then
      rm -rf ~/${runDirName}/dist
    fi

    tar -zxvf dist.tar.gz -C ~/${runDirName}
}

function deployPublicKey() {
    if [ -d ~/${runDirName}/public.pem ]; then
      rm -rf ~/${runDirName}/public.pem
    fi

    tar -zxvf dist.tar.gz -C ~/${runDirName}
    mv public.pem ~/${runDirName}
}

init_Dir
deployDist
deployPublicKey
#startTarget wc
startTarget rpt
#startTarget app
clean



