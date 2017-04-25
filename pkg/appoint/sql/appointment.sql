DROP TABLE IF EXISTS go_appoint_appointment;
DROP TABLE IF EXISTS go_appoint_order;
DROP TABLE IF EXISTS go_appoint_comment;
DROP TABLE IF EXISTS go_appoint_plan;
DROP TABLE IF EXISTS go_appoint_capacity_records;
DROP TABLE IF EXISTS go_appoint_sale_records;

CREATE TABLE go_appoint_appointment
(
  id VARCHAR(30)  primary key,
  appointtime bigint,
  org_code VARCHAR(30) references go_appoint_organization(org_code),
  planid VARCHAR(30) references go_appoint_plan(id),
  cardtype VARCHAR(10) not null,
  cardno VARCHAR(20) not null,
  mobile VARCHAR(15) not null,
  appointor VARCHAR(15) not null,
  merrystatus VARCHAR(10),
  status VARCHAR(10) not null,
  appoint_channel VARCHAR(30),
  company VARCHAR(50),
  "group" VARCHAR(50),
  remark VARCHAR(100),
  operator VARCHAR(15),
  operatetime bigint,
  orderid VARCHAR(30),
  commentid VARCHAR(30),
  appointednum integer,
  ifsingle boolean,
  ifcancel boolean
);

CREATE TABLE go_appoint_order(
  id VARCHAR(30)  primary key
);

CREATE TABLE go_appoint_comment(
    id VARCHAR(30) primary key,
    environment FLOAT,
    attitude FLOAT,
    breakfase FLOAT,
    details VARCHAR(500)
);

CREATE TABLE go_appoint_plan(
    id VARCHAR(30) primary key,
    name VARCHAR(30),
    avatar_img VARCHAR(50),
    detail_img VARCHAR(),
    checkups VARCHAR(30)[],
    ifshow boolean
);

CREATE TABLE go_appoint_capacity_records(
    org_code VARCHAR(30) references go_appoint_organization(org_code),
    date VARCHAR(10),
    used INTEGER DEFAULT 0
);

CREATE TABLE go_appoint_sale_records(
    org_code VARCHAR(30) references go_appoint_organization(org_code),
    sale_code VARCHAR(30),
    date VARCHAR(10),
    used INTEGER DEFAULT 0
);