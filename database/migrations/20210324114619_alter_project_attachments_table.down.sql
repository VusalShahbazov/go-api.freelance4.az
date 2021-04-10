ALTER TABLE  project_attachments
DROP FOREIGN  KEY fk_project_id2;

ALTER TABLE project_attachments
DROP COLUMN project_id;
