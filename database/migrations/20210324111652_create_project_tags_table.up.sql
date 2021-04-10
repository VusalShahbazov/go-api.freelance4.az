CREATE TABLE project_tags
(
    id              BIGINT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    tag_id          BIGINT(11) UNSIGNED,
    project_id      BIGINT(11) UNSIGNED NOT NULL,
    custom_tag_name VARCHAR(255),
    CONSTRAINT fk_projects_id1 FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE cascade,
    CONSTRAINT fk_tags_id1 FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE set null,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP NULL DEFAULT NULL
);