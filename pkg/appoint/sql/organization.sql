DROP TABLE IF EXISTS go_appoint_organization_basic_con;
DROP TABLE IF EXISTS go_appoint_organization_special_con;
DROP TABLE IF EXISTS go_appoint_organization CASCADE;

CREATE TABLE go_appoint_organization(  --分院表
   ORG_CODE         VARCHAR(30) PRIMARY KEY, --分院code
   ID               VARCHAR(30), --分院id
   NAME             VARCHAR(50), --分院名称
   imageUrl         VARCHAR(50), --分院简介 图片地址
   detailsUrl       VARCHAR(50), --分院详解 图片地址
   phone            VARCHAR(15), --分院客服电话
   DELETED          BOOLEAN DEFAULT FALSE --是否删除
);

CREATE TABLE go_appoint_organization_basic_con( --分院配置表
   ORG_CODE      VARCHAR(30) UNIQUE REFERENCES go_appoint_organization(ORG_CODE), --关联go_appoint_organization org
   CAPACITY      INTEGER DEFAULT 0, --分院容量
   WARNNUM       INTEGER DEFAULT 0, --分院提醒容量
   OFFDAYS       VARCHAR(30)[] DEFAULT ARRAY[]::VARCHAR[], --休息日
   AVOIDNUMBERS  INTEGER[] DEFAULT ARRAY[]::INTEGER[], --不使用预约号
   ip_address VARCHAR --分院所对的pinto服务地址
);

CREATE TABLE go_appoint_organization_special_con( --特殊项目限制表
   ORG_CODE         VARCHAR(30) NOT NULL REFERENCES go_appoint_organization(ORG_CODE), --关联go_appoint_organization org
   checkup_code     VARCHAR(30), --特殊项目的code
   CAPACITY          INTEGER --容量
);
