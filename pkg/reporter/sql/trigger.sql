DROP TABLE IF EXISTS go_togen_report;

CREATE TABLE go_togen_report (
   ID SERIAL PRIMARY KEY,
   EX_NO  VARCHAR(30)  NOT NULL,
   CREATETIME    timestamp(0) without time zone,
   FINISHTIME    timestamp(0) without time zone
);

CREATE OR REPLACE FUNCTION generateReport() RETURNS trigger AS $$
    BEGIN
        IF (TG_OP = 'INSERT') THEN
            perform genAllData(NEW.EX_NO);
        END IF;
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS gen_reporter_audit ON go_togen_report;
CREATE TRIGGER gen_reporter_audit AFTER INSERT ON go_togen_report
    FOR EACH ROW EXECUTE PROCEDURE generateReport();

CREATE OR REPLACE FUNCTION recordNewReporter() RETURNS trigger AS $$
    BEGIN
        IF NEW.status = 1080 THEN
            INSERT INTO go_togen_report(EX_NO, CREATETIME) Values(NEW.examination_no, current_timestamp);
        END IF;
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS examination_audit ON examination;

CREATE TRIGGER examination_audit AFTER INSERT OR UPDATE ON examination
    FOR EACH ROW EXECUTE PROCEDURE recordNewReporter();