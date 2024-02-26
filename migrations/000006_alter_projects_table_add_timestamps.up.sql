ALTER TABLE
    projects
ADD
    created_at TIMESTAMP NOT NULL DEFAULT(NOW());

ALTER TABLE
    projects
ADD
    updated_at TIMESTAMP NOT NULL DEFAULT(NOW());

ALTER TABLE
    projects
ADD
    deleted_at TIMESTAMP;
