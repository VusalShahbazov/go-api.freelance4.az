CREATE TABLE marked_projects
(
    id               BIGINT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_ud          BIGINT(11) UNSIGNED NOT NULL,
    project_id       BIGINT(11) UNSIGNED NOT NULL,
    CONSTRAINT fk_user_id6 FOREIGN KEY (user_ud) REFERENCES users (id) ON DELETE cascade,
    CONSTRAINT fk_project_id6 FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE cascade,
    created_at       TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);