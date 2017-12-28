-- +goose Up
INSERT INTO manpower.industry( type_id, industry)
VALUES (1, 'Mechanical'),
  (2, 'Electrical'),
  (3, 'Civil');

-- +goose Down
DELETE FROM manpower.industry;