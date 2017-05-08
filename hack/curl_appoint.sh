#!/usr/bin/env bash

target=192.168.199.198:9200

curl -XPOST ${target}/api/appointment -d '
{"planid":"1",
"appoint_time":1494497277,
"org_code":"0001002",
"cardno":"210221198906068120",
"cardtype":"身份证",
"mobile":"17744524309",
"appointor":"小潘潘的测试",
"merrystatus":"未婚",
"status":"预约",
"appoint_channel":"微信",
"company":"","group":"","remark":"",
"operator":"operator13",
"operate_time":1494497277,
"ifsingle":false,
"ifcancel":false
}'




