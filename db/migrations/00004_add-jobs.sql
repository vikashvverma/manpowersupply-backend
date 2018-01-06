-- +goose Up
INSERT INTO manpower.job(job_id, title, industry, location, date_created, date_updated, available, type_id)
VALUES (1001, 'Supervisor', 'Mechanical', 'Surat', now(), now(), true, 1),
  (2002, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (2003, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (2004, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (2005, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (2006, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (2007, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (2008, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (2009, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (2010, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (2011, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (2012, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (2013, 'Worker', 'Electrical', 'Delhi', now(), now(), true, 2),
  (3001, 'Contractor', 'Civil', 'Delhi', now(), now(), true, 3),
  (3002, 'Contractor', 'Civil', 'Delhi', now(), now(), true, 3),
  (3003, 'Contractor', 'Civil', 'Delhi', now(), now(), true, 3),
  (3004, 'Contractor', 'Civil', 'Delhi', now(), now(), true, 3),
  (3005, 'Contractor', 'Civil', 'Delhi', now(), now(), true, 3),
  (3006, 'Contractor', 'Civil', 'Delhi', now(), now(), true, 3),
  (3007, 'Contractor', 'Civil', 'Delhi', now(), now(), true, 3),
  (3008, 'Contractor', 'Civil', 'Delhi', now(), now(), true, 3),
  (3009, 'Contractor', 'Civil', 'Delhi', now(), now(), true, 3);

-- +goose Down
DELETE FROM manpower.job;