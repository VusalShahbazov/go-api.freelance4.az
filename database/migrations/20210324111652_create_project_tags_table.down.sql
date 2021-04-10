ALTER TABLE project_tags
DROP FOREIGN KEY fk_projects_id1;
ALTER TABLE project_tags
DROP FOREIGN KEY fk_tags_id1;


drop table if exists project_tags;