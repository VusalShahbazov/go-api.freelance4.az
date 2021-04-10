ALTER TABLE comments
DROP
FOREIGN KEY fk_user_id4;

ALTER TABLE comments
DROP
FOREIGN KEY fk_project_id4;



DROP TABLE IF EXISTS  comments;