ALTER TABLE  projects
DROP FOREIGN  KEY fk_category_id1;

ALTER TABLE projects
DROP COLUMN category_id;
