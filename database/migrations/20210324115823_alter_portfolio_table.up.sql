ALTER TABLE portfolio
    ADD COLUMN user_id BIGINT(11) UNSIGNED;

ALTER TABLE portfolio
    ADD CONSTRAINT  fk_user_id1
        FOREIGN KEY (user_id) REFERENCES users(id)