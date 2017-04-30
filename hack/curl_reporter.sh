#!/usr/bin/env bash


curl 192.168.199.198:9100/api/report/list

curl  192.168.199.198:9100/api/report?examination_no=0001160001912
curl  127.0.0.1:9100/api/report?examination_no=0001160005705
curl  10.1.0.235:9100/api/report?examination_no=0001160005705

curl  192.168.199.198:9100/api/report/list?username=cnadmin&password=202cb962ac59075b964b07152d234b70

curl  -XPOST '127.0.0.1:9100/api/report/status?username=cnadmin&password=202cb962ac59075b964b07152d234b70&examination_no=0001160007406&status=printed'

