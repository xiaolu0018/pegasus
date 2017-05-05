#!/usr/bin/env bash

function makeDir() {
   if [ -d ~/$1 ]; then
      rm -rf ~/$1
   fi

    mkdir ~/$1
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
    makeDir pegasus_$1
    cp pegasus ~/pegasus_$1/pegasus_$1
    deployDist pegasus_$1
}

function target::Start() {
    cd ~/pegasus_$1
    ./pegasus_$1 $1 $2> ./$1_$2.log 2>&1 &
}

function startTarget() {
    cd ~
    target::Kill_old $1
    target::Prepare $1
    target::Start $1 $2
}

function deployDist() {
    if [ -d ~/$1/dist ]; then
      rm -rf ~/$1/dist
    fi

    tar -zxvf dist.tar.gz -C ~/$1
}

#function deployPublicKey() {
#    if [ -d ~/${runDirName}/public.pem ]; then
#      rm -rf ~/${runDirName}/public.pem
#    fi
#
#    tar -zxvf dist.tar.gz -C ~/${runDirName}
#    mv public.pem ~/${runDirName}
#}


startTarget wc start
startTarget wc start-activity
startTarget rpt start
startTarget app start

clean



