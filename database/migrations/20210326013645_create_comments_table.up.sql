CREATE TABLE comments
(
    id               BIGINT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_ud          BIGINT(11) UNSIGNED NOT NULL,
    project_id       BIGINT(11) UNSIGNED NOT NULL,
    description      TEXT    NOT NULL,
    attach_portfolio boolean not null default 0,
    pinned           boolean not null default 0,
    CONSTRAINT fk_user_id4 FOREIGN KEY (user_ud) REFERENCES users (id) ON DELETE cascade,
    CONSTRAINT fk_project_id4 FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE cascade,
    created_at       TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at       DATETIME         DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at       TIMESTAMP NULL DEFAULT NULL
);