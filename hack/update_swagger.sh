#!/usr/bin/env bash

filepath=$(cd "$(dirname "$0")"; pwd)

scp $filepath/../docs/api.json weixin@192.168.199.198:/home/weixin/tomcat/webapps/swagger-ui/dist
