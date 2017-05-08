DROP TABLE IF EXISTS go_appoint_login_user;

CREATE TABLE go_appoint_login_user(
    loginaccount VARCHAR(30) PRIMARY KEY, --登录账号
    password VARCHAR(30), --密码
    loginname VARCHAR(20), --登录名
    orgCode VARCHAR(20) --所属分院
);