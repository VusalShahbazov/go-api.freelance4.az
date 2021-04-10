ALTER TABLE projects
    DROP COLUMN  is_active;



ALTER TABLE projects
    DROP FOREIGN KEY fk_creator_id1;


ALTER TABLE projects
    DROP FOREIGN KEY fk_freelancer_id1;

ALTER TABLE projects
    DROP COLUMN  freelancer_id;


ALTER TABLE projects
    DROP COLUMN  creator_id;