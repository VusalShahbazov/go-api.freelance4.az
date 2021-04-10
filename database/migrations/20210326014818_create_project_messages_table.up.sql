CREATE TABLE project_messages
(
    id               BIGINT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_ud          BIGINT(11) UNSIGNED NOT NULL,
    project_id       BIGINT(11) UNSIGNED NOT NULL,
    body             VARCHAR(255)   NOT NULL,
    is_read          boolean default 0,
    CONSTRAINT fk_user_id3 FOREIGN KEY (user_ud) REFERENCES users (id) ON DELETE cascade,
    CONSTRAINT fk_project_id3 FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE cascade,
    created_at       TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at       DATETIME         DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at       TIMESTAMP NULL DEFAULT NULL
);