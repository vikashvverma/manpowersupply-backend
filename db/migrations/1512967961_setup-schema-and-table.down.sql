ALTER TABLE manpower.job DROP CONSTRAINT job_fk1;

ALTER TABLE manpower.job DROP CONSTRAINT job_fk0;

ALTER TABLE manpower.query DROP CONSTRAINT query_fk0;

DROP TABLE manpower.job;

DROP TABLE manpower.job_type;

DROP TABLE manpower.query;

DROP TABLE manpower.party;

DROP SCHEMA manpower;