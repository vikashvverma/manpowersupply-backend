CREATE SCHEMA manpower;

CREATE TABLE manpower.party (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	address VARCHAR(100) NOT NULL,
	phone varchar(15) NOT NULL,
	mobile varchar(15) NOT NULL,
	email varchar(40) NOT NULL
);



CREATE TABLE manpower.query (
	id SERIAL,
	queryer_id INT NOT NULL,
	query varchar(400) NOT NULL,
	query_date timestamptz  NOT NULL DEFAULT now(),
	PRIMARY KEY (id)
);

CREATE TABLE manpower.job_type (
	id SERIAL,
	type_id INT NOT NULL UNIQUE,
	industry varchar(10) NOT NULL UNIQUE,
	PRIMARY KEY (id)
);

CREATE TABLE manpower.job (
	id SERIAL,
	job_id INT NOT NULL UNIQUE,
	title varchar(10) NOT NULL,
	industry varchar(10) NOT NULL,
	location varchar(10) NOT NULL,
	date_created timestamptz NOT NULL,
	date_updated timestamptz NOT NULL,
	available BOOLEAN NOT NULL,
	type_id INT NOT NULL,
	PRIMARY KEY (id)
);



ALTER TABLE manpower.query ADD CONSTRAINT query_fk0 FOREIGN KEY (queryer_id) REFERENCES manpower.party(id);

ALTER TABLE manpower.job ADD CONSTRAINT job_fk0 FOREIGN KEY (industry) REFERENCES manpower.job_type(industry);

ALTER TABLE manpower.job ADD CONSTRAINT job_fk1 FOREIGN KEY (type_id) REFERENCES manpower.job_type(type_id);

