#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

tar -czvf dist.tar.gz dist

scp pegasus public.pem hack/deploy_server.sh dist.tar.gz weixin@192.168.199.198:~

target="cd ~ && chmod +x deploy_server.sh && ./deploy_server.sh $1"
ssh weixin@192.168.199.198 $target

