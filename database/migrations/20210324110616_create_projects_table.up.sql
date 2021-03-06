CREATE TABLE projects
(
    id          BIGINT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name        VARCHAR(255)   NOT NULL,
    slug        VARCHAR(255)   NOT NULL UNIQUE,
    description TEXT           NOT NULL,
    price       DECIMAL(10, 2) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP NULL DEFAULT NULL
);