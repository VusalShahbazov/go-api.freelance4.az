ALTER TABLE projects
    ADD COLUMN category_id  BIGINT(11) UNSIGNED ;

ALTER TABLE projects
    ADD CONSTRAINT  fk_category_id1
        FOREIGN KEY (category_id) REFERENCES project_categories(id);