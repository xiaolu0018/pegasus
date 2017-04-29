DROP TABLE IF EXISTS go_appoint_login_user;

CREATE TABLE go_appoint_login_user(
    loginaccount VARCHAR(30) PRIMARY KEY,
    password VARCHAR(30),
    loginname VARCHAR(20),
    orgcode VARCHAR(20)
);