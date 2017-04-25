#!/usr/bin/env bash

target=127.0.0.1:9200


curl -v -XPOST ${target}/api/appointment -d '{"appointor":"潘新元","cardtype":"身份证", "cardno":"210221198906068111", "mobile":"17777777777", "merrystatus":"未婚", "org_code":"xxx"}'

#curl ${target}/api/organazations
#
#curl -XPUT ${target}/api/organazation/000100102/config/basic -d '{"capacity":10,"warnnum":9,"offdays":["2017-04-28","2017-05-03"],"avoid_numbers":[13,4]}'
#
#curl -XPOST  ${target}/api/organazation/000100102/config/special '{"sale_code":"0004","capacity":8}'
#
#curl -XPOST ${target}/api/appointment/appoint13/cancel
#
#curl -XPOST ${target}/api/appointment -d '{"planid":"1","appointtime":1493029303,"org_code":"000101","cardno":"cardid1","cardtype":"cardType1","mobile":"mobile1","appointor":"httptext","merrystatus":"未婚","status":"预约","appoint_channel":"微信","channel_appointor_id":"","company":"","group":"","remark":"","operator":"operator13","operatetime":1492856503,"orderid":"order13","commentid":"","appointednum":0,"ifsingle":false,"ifcancel":false}'
#
#curl -XPUT ${target}/api/appointment -d '{"id":"112313dfdfdf","planid":"1","appointtime":1493115913,"org_code":"000101","cardno":"cardid1","cardtype":"cardType1","mobile":"mobile1","appointor":"httptext","merrystatus":"未婚","status":"预约","appoint_channel":"微信","channel_appointor_id":"","company":"","group":"","remark":"","operator":"operator13","operatetime":1492856503,"orderid":"order13","commentid":"","appointednum":0,"ifsingle":false,"ifcancel":false}'
#
#curl ${target}/api/appointment/112313dfdfdf
#
#curl ${target}/api/appointmenlist
#
#curl -XPOST ${target}/api/appointment/112313/comment -d '{"environment":0,"attitude":0,"breakfast":0,"details":""}'





