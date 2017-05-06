DROP TABLE IF EXISTS go_appoint_appointment;
DROP TABLE IF EXISTS go_appoint_order;
DROP TABLE IF EXISTS go_appoint_comment;
DROP TABLE IF EXISTS go_appoint_plan;
DROP TABLE IF EXISTS go_appoint_capacity_records;
DROP TABLE IF EXISTS go_appoint_sale_records;
DROP TABLE IF EXISTS go_appoint_banner;

CREATE TABLE go_appoint_appointment  --appoint 服务 预约数据表
(
  id VARCHAR(30)  primary key,
  appointtime bigint, --预约体检时间
  org_code VARCHAR(30) references go_appoint_organization(org_code),  --分院代码
  planid VARCHAR(30) references go_appoint_plan(id),  --套餐id
  sale_codes VARCHAR(30)[], --销售项目code
  cardtype VARCHAR(10) not null, --证件类型
  cardno VARCHAR(20) not null, --证件号
  mobile VARCHAR(15) not null, --手机号
  appointor VARCHAR(30) not null, --预约人姓名
  appointorid VARCHAR(30), --预约人id
  merrystatus VARCHAR(10), --婚姻状态
  status VARCHAR(10) not null, --体检状态
  appoint_channel VARCHAR(30), --预约渠道
  company VARCHAR(50), --公司名
  "group" VARCHAR(50), --分组名
  remark VARCHAR(100), --备注
  operator VARCHAR(15), --操作人
  operatetime bigint, --操作时间
  orderid VARCHAR(30), --订单id
  commentid VARCHAR(30), --评论id
  address VARCHAR(100), --地址
  appointednum integer, --体检号
  reportid VARCHAR(30), --报告id
  bookno VARCHAR(30), --pinto服务中的book_record的对应订阅号
  ifsingle boolean, --是否散客
  ifcancel boolean --是否取消
);

CREATE TABLE go_appoint_order( --订单表
  id VARCHAR(30)  primary key
);

CREATE TABLE go_appoint_comment( --预约评论表
    id VARCHAR(30) primary key,
    environment FLOAT, --环境评分
    attitude FLOAT, --态度评分
    breakfase FLOAT, --早餐评分
    details VARCHAR(500), --评论
    conclusion VARCHAR(10) --总评
);

CREATE TABLE go_appoint_plan( --套餐表
    id VARCHAR(30) primary key,
    name VARCHAR(30), --套餐名
    avatar_img VARCHAR(50), --套餐简介 图片地址
    detail_img VARCHAR(), --套餐详解 图片地址
    sale_codes VARCHAR(30)[], --套餐所含的销售项
    ifshow boolean --是否展示
);

CREATE TABLE go_appoint_banner( --微信banner表
    id VARCHAR(30) primary key,
    pos INTEGER, --位置
    imageurl VARCHAR(30), --简介 图片地址
    redirecturl VARCHAR(30), --详解 图片地址
    ifshow boolean --是否显示
);

CREATE TABLE go_appoint_capacity_records( --分院预约记录表
    org_code VARCHAR(30) references go_appoint_organization(org_code), --外键org_code
    date VARCHAR(10), --日期
    used INTEGER DEFAULT 0 --已预约人数
);

CREATE TABLE go_appoint_sale_records( --特殊项目预约限制记录
    org_code VARCHAR(30) references go_appoint_organization(org_code),--外键org_code
    checkup_code VARCHAR(30), --特殊项目code
    date VARCHAR(10), --日期
    used INTEGER DEFAULT 0 --已预约人数
);