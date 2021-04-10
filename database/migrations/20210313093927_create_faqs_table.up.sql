CREATE TABLE faqs
(
    id          BIGINT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    answer_az   MEDIUMTEXT not null,
    answer_ru   MEDIUMTEXT not null,
    question_az MEDIUMTEXT not null,
    question_ru MEDIUMTEXT not null,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP NULL DEFAULT NULL
);