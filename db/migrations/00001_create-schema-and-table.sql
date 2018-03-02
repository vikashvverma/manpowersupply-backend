-- +goose Up
CREATE SCHEMA manpower;

CREATE TABLE manpower.party (
  id      SERIAL PRIMARY KEY,
  name    VARCHAR(50)  NOT NULL,
  address VARCHAR(100) NOT NULL,
  company VARCHAR(100),
  website VARCHAR(100),
  city    VARCHAR(50),
  state   VARCHAR(50),
  pin     VARCHAR(15),
  country VARCHAR(15)  NOT NULL,
  mobile  VARCHAR(15)  NOT NULL,
  email   VARCHAR(40),
  phone   VARCHAR(40)
);


CREATE TABLE manpower.query (
  id         SERIAL,
  queryer_id INT         NOT NULL,
  query      VARCHAR(400),
  industry   VARCHAR(20),
  title      VARCHAR(50),
  query_date TIMESTAMPTZ NOT NULL DEFAULT now(),
  PRIMARY KEY (id)
);

CREATE TABLE manpower.industry (
  id       SERIAL,
  type_id  INT         NOT NULL UNIQUE,
  industry VARCHAR(20) NOT NULL UNIQUE,
  PRIMARY KEY (id)
);

CREATE TABLE manpower.job_type (
  id      SERIAL,
  type_id INT         NOT NULL,
  title   VARCHAR(50) NOT NULL,
  UNIQUE (type_id, title),
  PRIMARY KEY (id)
);

CREATE TABLE manpower.job (
  id           SERIAL,
  job_id       INT         NOT NULL UNIQUE,
  title        VARCHAR(20) NOT NULL,
  industry     VARCHAR(20) NOT NULL,
  location     VARCHAR(20) NOT NULL,
  date_created TIMESTAMPTZ NOT NULL DEFAULT now(),
  date_updated TIMESTAMPTZ NOT NULL DEFAULT now(),
  available    BOOLEAN     NOT NULL,
  type_id      INT         NOT NULL,
  PRIMARY KEY (id)
);


ALTER TABLE manpower.query
  ADD CONSTRAINT query_fk0 FOREIGN KEY (queryer_id) REFERENCES manpower.party (id);

ALTER TABLE manpower.job
  ADD CONSTRAINT job_fk0 FOREIGN KEY (industry) REFERENCES manpower.industry (industry);

ALTER TABLE manpower.job
  ADD CONSTRAINT job_fk1 FOREIGN KEY (type_id) REFERENCES manpower.industry (type_id);

-- +goose Down
ALTER TABLE manpower.job
  DROP CONSTRAINT job_fk1;

ALTER TABLE manpower.job
  DROP CONSTRAINT job_fk0;

ALTER TABLE manpower.query
  DROP CONSTRAINT query_fk0;

DROP TABLE manpower.job;

DROP TABLE manpower.job_type;

DROP TABLE manpower.industry;

DROP TABLE manpower.query;

DROP TABLE manpower.party;

DROP SCHEMA manpower;
