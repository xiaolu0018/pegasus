#!/usr/bin/env bash

curl 192.168.199.198:9200/api/organazations

curl -XPUT 192.168.199.198:9200/api/organazation/000100102/config/basic -d '{"capacity":10,"warnnum":9,"offdays":["2017-04-28","2017-05-03"],"avoid_numbers":[13,4]}'

curl -XPOST  192.168.199.198:9200/api/organazation/000100102/config/special '{"sale_code":"0004","capacity":8}'

curl -XPOST 192.168.199.198:9200/api/appointment/appoint13/cancel

curl -XPOST 192.168.199.198:9200/api/appointment -d '{"planid":"1","appointtime":1493029303,"orgcode":"000101","cardno":"cardid1","cardtype":"cardType1","mobile":"mobile1","appointor":"httptext","merrystatus":"未婚","status":"预约","appoint_channel":"微信","channel_appointor_id":"","company":"","group":"","remark":"","operator":"operator13","operatetime":1492856503,"orderid":"order13","commentid":"","appointednum":0,"ifsingle":false,"ifcancel":false}'

curl -XPUT 192.168.199.198:9200/api/appointment -d '{"id":"112313dfdfdf","planid":"1","appointtime":1493115913,"orgcode":"000101","cardno":"cardid1","cardtype":"cardType1","mobile":"mobile1","appointor":"httptext","merrystatus":"未婚","status":"预约","appoint_channel":"微信","channel_appointor_id":"","company":"","group":"","remark":"","operator":"operator13","operatetime":1492856503,"orderid":"order13","commentid":"","appointednum":0,"ifsingle":false,"ifcancel":false}'

curl 192.168.199.198:9200/api/appointment/112313dfdfdf

curl 192.168.199.198:9200/api/appointmenlist

curl -XPOST 192.168.199.198:9200/api/appointment/112313/comment -d '{"environment":0,"attitude":0,"breakfast":0,"details":""}'

