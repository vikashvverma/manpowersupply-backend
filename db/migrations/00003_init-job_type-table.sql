-- +goose Up
INSERT INTO manpower.job_type(type_id, title)
VALUES (1, 'Engineer'),
  (1, 'Supervisor'),
  (1, 'Piping Fabricator'),
  (1, 'Piping Fitter'),
  (1, 'Welder'),
  (1, 'Grinder'),
  (1, 'Gas Cutter'),
  (1, 'Rigger'),
  (1, 'Khalasi'),
  (1, 'Helper'),
  (2, 'Supervisor'),
  (2, 'Electrician'),
  (2, 'Fitter'),
  (2, 'Rigger'),
  (2, 'Khalasi'),
  (2, 'Helper'),
  (3, 'Engineer'),
  (3, 'Supervisor'),
  (3, 'Fitter'),
  (3, 'Carpenter'),
  (3, 'Bar Bender'),
  (3, 'Massion'),
  (3, 'Helper');
-- +goose Down
DELETE FROM manpower.job_type;