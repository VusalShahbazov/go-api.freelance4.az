ALTER TABLE project_attachments
    ADD COLUMN project_id BIGINT(11) UNSIGNED;

ALTER TABLE project_attachments
    ADD CONSTRAINT fk_project_id2
        FOREIGN KEY (project_id) REFERENCES projects(id);