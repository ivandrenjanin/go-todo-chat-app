ALTER TABLE
    projects DROP COLUMN created_at;

ALTER TABLE
    projects
ADD
    updated_at;

ALTER TABLE
    projects
ADD
    deleted_at;
