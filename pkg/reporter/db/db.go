package db

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
)

//"postgres://postgres:postgresql2016@192.168.199.216:5432/pinto?sslmode=disable"
var readDB *sql.DB
var err error

func Init(user, passwd, ip, port, db string) error {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, ip, port, db)

	readDB, err = sql.Open("postgres", addr)
	return err
}

func GetDB() *sql.DB {
	return readDB
}

func InitFunction(user, passwd, ip, port, db string) error {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, ip, port, db)

	writeDB, err := sql.Open("postgres", addr)
	if err != nil {
		return err
	}
	defer writeDB.Close()

	_, err = writeDB.Exec(initCmd)
	return err
}

// **these sql functions need sync with sql/function.sql**
var initCmd = `
drop FUNCTION IF EXISTS  arrayToArrStr(arr text[]);
drop FUNCTION IF EXISTS  arrayToObjStr(arr text[]);
drop FUNCTION IF EXISTS  arrayToObjStr2(arr text[]);
drop FUNCTION IF EXISTS  arrayToArrStr2(arr text[]);
drop FUNCTION IF EXISTS  checkNull(s text);
drop FUNCTION IF EXISTS  getCheckupStr(exam_no varchar);
drop FUNCTION IF EXISTS  getSelectedStr(exam_no varchar);
drop FUNCTION IF EXISTS  getCheckAndItems(exam_no varchar);
drop FUNCTION IF EXISTS  getFinalDiagoseStr(exam_no varchar);
drop FUNCTION IF EXISTS  genFinalExam(exam_no varchar);
drop FUNCTION IF EXISTS  getItemStr(exam_no varchar, ck_code varchar);
drop FUNCTION IF EXISTS  getImageStr(exam_no varchar);
drop FUNCTION IF EXISTS  getSingles(exam_no varchar);
drop FUNCTION IF EXISTS  genAllData(exam_no varchar);
drop FUNCTION IF EXISTS  getReport(exam_no VARCHAR);

CREATE OR REPLACE FUNCTION arrayToArrStr(arr text[]) RETURNS text AS $$
   BEGIN
    return concat('[^[', array_to_string(arr, '^,^'), ']^]');
   END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION arrayToObjStr(arr text[]) RETURNS text AS $$
   BEGIN
    return concat('{^{', array_to_string(arr, '^:^'), '}^}');
   END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION arrayToArrStr2(arr text[]) RETURNS text AS $$
   BEGIN
    return concat('[%[', array_to_string(arr, '%,%'), ']%]');
   END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION arrayToObjStr2(arr text[]) RETURNS text AS $$
   BEGIN
    return concat('{%{', array_to_string(arr, '%:%'), '}%}');
   END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION checkNull(s text) RETURNS TEXT AS $$
    BEGIN
        IF (s = '') THEN
            RETURN 'NULL';
        ELSIF (s is NULL) THEN
            RETURN 'NULL';
        ELSE
            RETURN s;
        END IF;
    END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION getCheckupStr(exam_no varchar) RETURNS text AS $$
    DECLARE
        data record;
        ret  text[];
    BEGIN
        FOR data IN
        SELECT
            C.checkup_name, d.department_name, ec.checkup_status
        FROM examination T
        LEFT JOIN examination_checkup ec ON T .examination_no = ec.examination_no
        LEFT JOIN checkup C ON ec.checkup_code = C .checkup_code
        LEFT JOIN department d ON ec.department_code = d.department_code
        WHERE T .examination_no = exam_no AND C .is_valid = 1
        ORDER BY ec.checkup_status, d.department_code, C.order_position
   		LOOP
            select array_append(ret,  arrayToObjStr(ARRAY[data.checkup_name, data.department_name, data.checkup_status]::text[])) into ret;
   		END LOOP;

   		RETURN arrayToArrStr(ret);
    END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION getSelectedStr(exam_no varchar) RETURNS text AS $$
    DECLARE
        data record;
        selecteds text[];
    BEGIN
        FOR data IN
            SELECT P.selected_code
	        FROM personal_health_info P, examination b
	        WHERE P.person_code = b.person_code
	        AND b.examination_no = exam_no
	        ORDER BY P.person_code
        LOOP
            SELECT array_append(selecteds, data.selected_code::text) into selecteds;
        END LOOP;

        RETURN arrayToArrStr(selecteds);
    END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION getCheckAndItems(exam_no varchar) RETURNS text AS $$
    DECLARE
        data        record;
        tmp         text[];
    BEGIN
     FOR data IN
        SELECT
        DEP.department_name,
        I.item_name, EX_I.item_value, EX_I.exception_arrow, EX_I.reference_description, I.examination_unit,
        EX_CK.diagnose_result,
        DEP.doctor_sign, M.previous_name username, mm.previous_name,
        CK.checkup_type_code, CK.department_code
        FROM examination EX
        LEFT JOIN examination_checkup EX_CK ON EX.examination_no = EX_CK.examination_no
        LEFT JOIN checkup CK ON EX_CK.checkup_code = CK.checkup_code
        LEFT JOIN department DEP ON CK.department_code = DEP.department_code
        LEFT JOIN examination_item EX_I ON EX.examination_no = EX_I.examination_no AND EX_CK.checkup_code = EX_I.checkup_code
        LEFT JOIN item I ON EX_I.item_code = I.item_code
        LEFT JOIN manager M ON EX_CK.diagnose_manager_code = M.manager_code
        LEFT JOIN MANAGER mm ON EX_CK.check_manager_code = mm.manager_code
        WHERE EX_CK.checkup_code IS NOT NULL
        and CK.checkup_type_code IN ('0', '1')
        AND EX_CK.checkup_status = 2
        AND EX.examination_no = exam_no
        AND (
	        (
		        I.validate_type = 1
		        AND EX_I.item_value IS NOT NULL
	        )
	        OR I.validate_type = 0
        )
        AND EX_CK.department_code NOT IN ('63')
        ORDER BY DEP.department_code, ck.checkup_code, i.order_position, I.item_code
        LOOP
            select array_append(tmp, arrayToObjStr(ARRAY[checkNull(data.department_name), checkNull(data.item_name),
            checkNull(data.item_value), checkNull(data.exception_arrow), checkNull(data.reference_description),
            checkNull(data.examination_unit), checkNull(data.diagnose_result), checkNull(data.doctor_sign),
            checkNull(data.username), checkNull(data.previous_name),
            checkNull(data.checkup_type_code), checkNull(data.department_code)]::text[])) into tmp;
        END LOOP;

        return arrayToArrStr(tmp);

    END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION getFinalDiagoseStr(exam_no varchar) RETURNS text AS $$
    DECLARE
        data        record;
        tmp         text[];
    BEGIN
    FOR DATA IN
        SELECT
                C .checkup_name,
                T .analyzse_doctor,
        		T .analyse_advice,
				ec.diagnose_result

			FROM
				examination_analyse T
			LEFT JOIN checkup C ON T .checkup_code = C .checkup_code
			LEFT JOIN examination_checkup ec ON (
				T .checkup_code = ec.checkup_code
				AND T .examination_no = ec.examination_no
			)
			WHERE
				T .examination_no = exam_no
			AND T .department_code <> '63'
			AND (
				analyse_advice IS NOT NULL
				AND analyse_advice <> ' '
				AND analyse_advice <> ''
			)
			ORDER BY
				C .order_position
    LOOP
        SELECT array_append(tmp, arrayToObjStr(ARRAY[data.checkup_name, data.analyzse_doctor, data.analyse_advice, data.diagnose_result]::text[]))
        INTO tmp;
    END LOOP;

    return arrayToArrStr(tmp);
    END;

$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION genFinalExam(exam_no varchar) RETURNS text AS $$
    DECLARE
        t1      text;
        t2      text;
        t3      text;
    BEGIN
    SELECT SUBSTR(updatetime, 0, 11), finalexamination FROM examination_analyse_finalexamination
    WHERE examination_no = exam_no into t1, t2;
    SELECT key_value FROM public.con_global_config where key_name='inspection_doctor' into t3;

    RETURN arrayToObjStr(ARRAY[t1, t2, t3]);
    END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION getItemStr(exam_no varchar, ck_code varchar) RETURNS text AS $$
    DECLARE
        data record;
        tmp text[];
    BEGIN
        FOR data IN
             select item_name, item_value from examination_item where checkup_code = ck_code and examination_no = exam_no
             LOOP
                SELECT array_append(tmp, arrayToObjStr2(ARRAY[data.item_name, data.item_value])) into tmp;
             END LOOP;
        RETURN arrayToArrStr2(tmp);
    END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION getImageStr(exam_no varchar) RETURNS text AS $$
    DECLARE
        data record;

        tmp text[];
    BEGIN
        FOR data IN
            SELECT ec.checkup_code, ec.diagnose_result, ec.image_url, c.checkup_name,
            d.doctor_sign, m.previous_name
            FROM examination_checkup ec
            LEFT JOIN checkup c ON ec.checkup_code = c.checkup_code
            LEFT JOIN manager m ON ec.diagnose_manager_code=m.manager_code
            LEFT JOIN department d ON d.department_code = ec.department_code
            WHERE ec.examination_no = exam_no and ec.image_url is not null
            AND EC.checkup_code NOT IN (
            SELECT regexp_split_to_table(key_value,',') from con_global_config where  key_name = 'not_print_code' OR  key_name = 'single_print_code'
            )

        LOOP
            select array_append(tmp, arrayToObjStr(ARRAY[data.checkup_name, data.diagnose_result, data.image_url, getItemStr(exam_no, data.checkup_code), data.doctor_sign, data.previous_name])) into tmp;
        END LOOP;

        RETURN arrayToArrStr(tmp);
    END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION getSingles(exam_no varchar) RETURNS text AS $$
    DECLARE
        data record;
        tmp text[];
    BEGIN
        FOR data IN
        SELECT image_url
        FROM examination_imageinformation T
        WHERE T.examination_no = exam_no
        AND T.image_url IS NOT NULL
        AND T.checkup_code IN ( SELECT regexp_split_to_table(key_value, ',') FROM con_global_config WHERE key_name = 'single_print_code')
        ORDER BY  checkup_code,t.createtime
    LOOP
        SELECT array_append(tmp, arrayToObjStr(ARRAY[data.image_url])) into tmp;
        END LOOP;
        RETURN arrayToArrStr(tmp);
    END;
$$ LANGUAGE plpgsql;
`
