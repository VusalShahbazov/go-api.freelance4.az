ALTER TABLE projects
    ADD COLUMN is_active boolean not null default 1;


ALTER TABLE projects
    ADD COLUMN creator_id BIGINT(11) UNSIGNED NOT NULL;


ALTER TABLE projects
    ADD CONSTRAINT fk_creator_id1 FOREIGN  KEY  (creator_id) REFERENCES users(id);


ALTER TABLE projects
    ADD COLUMN freelancer_id BIGINT(11) UNSIGNED;


ALTER TABLE projects
    ADD CONSTRAINT fk_freelancer_id1 FOREIGN  KEY  (freelancer_id) REFERENCES users(id);
