DROP TABLE IF EXISTS go_weixin_user;
DROP TABLE IF EXISTS go_weixin_user_health;

CREATE TABLE go_weixin_user
(
  id VARCHAR(30)  primary key,
  openid    VARCHAR(30) UNIQUE,
  cardtype VARCHAR(10) ,
  cardno VARCHAR(20),
  mobile VARCHAR(15) ,
  name VARCHAR(10) ,
  sex VARCHAR(2),
  merrystatus VARCHAR(10) ,
  address_province VARCHAR(10),
  address_city VARCHAR(10),
  address_district VARCHAR(10),
  address_details VARCHAR(30),
  ifonlyneed_electronic_report boolean DEFAULT FALSE,

  wc_nickname VARCHAR(100),
  wc_sex VARCHAR (2),
  wc_province VARCHAR (10),
  wc_city VARCHAR (20),
  wc_country VARCHAR(10),
  wc_headimgurl VARCHAR (200),
  healthid VARCHAR(30)
);

CREATE TABLE go_weixin_user_health
(
    id VARCHAR(30) primary key,
    past_history VARCHAR(10)[] DEFAULT ARRAY[]::VARCHAR[],
    family_medical_history VARCHAR(10)[] DEFAULT ARRAY[]::VARCHAR[],
    exam_frequency VARCHAR(10)[] DEFAULT ARRAY[]::VARCHAR[],
    past_exam_exception VARCHAR(10)[] DEFAULT ARRAY[]::VARCHAR[],
    psychological_pressure VARCHAR(10)[] DEFAULT ARRAY[]::VARCHAR[],
    food_habits VARCHAR(10)[] DEFAULT ARRAY[]::VARCHAR[],
    eating_habits VARCHAR(10)[] DEFAULT ARRAY[]::VARCHAR[],
    drink_habits VARCHAR(10)[] DEFAULT ARRAY[]::VARCHAR[],
    smoke_history VARCHAR(10)[] DEFAULT ARRAY[]::VARCHAR[]
);