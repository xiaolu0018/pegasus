DROP TABLE IF EXISTS go_appoint_organization_basic_con;
DROP TABLE IF EXISTS go_appoint_organization_special_con;
DROP TABLE IF EXISTS go_appoint_organization CASCADE;

CREATE TABLE go_appoint_organization(
   ORG_CODE         VARCHAR(30) PRIMARY KEY,
   ID               VARCHAR(30),
   NAME             VARCHAR(50),
   imageUrl         VARCHAR(50),
   detailsUrl       VARCHAR(50),
   phone            VARCHAR(15),
   DELETED          BOOLEAN DEFAULT FALSE
);

CREATE TABLE go_appoint_organization_basic_con(
   ORG_CODE      VARCHAR(30) UNIQUE REFERENCES go_appoint_organization(ORG_CODE),
   CAPACITY      INTEGER DEFAULT 0,
   WARNNUM       INTEGER DEFAULT 0,
   OFFDAYS       VARCHAR(10)[] DEFAULT ARRAY[]::VARCHAR[],
   AVOIDNUMBERS  INTEGER[] DEFAULT ARRAY[]::INTEGER[],
   ip_address VARCHAR
);

CREATE TABLE go_appoint_organization_special_con(
   ORG_CODE      VARCHAR(30) NOT NULL REFERENCES go_appoint_organization(ORG_CODE),
   SALE_CODE     VARCHAR(30),
   CAPACITY      INTEGER
);
