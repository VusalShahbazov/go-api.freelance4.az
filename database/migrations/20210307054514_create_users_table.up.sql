CREATE TABLE users
(
    id            BIGINT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    last_name     VARCHAR(255) not null,
    first_name    VARCHAR(255) not null,
    email         VARCHAR(255) not null unique,
    password      VARCHAR(255) not null,
    is_freelancer BOOLEAN   DEFAULT 1,
    phone_number  VARCHAR(255),
    avatar        VARCHAR(255),
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP NULL DEFAULT NULL,
    INDEX         email_and_deleted_at (email,deleted_at)
);