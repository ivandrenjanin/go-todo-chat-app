ALTER TABLE
    users
ALTER COLUMN
    "first_name" TYPE VARCHAR(50);

ALTER TABLE
    users
ALTER COLUMN
    "last_name" TYPE VARCHAR(50);

ALTER TABLE
    users
ALTER COLUMN
    "email" TYPE VARCHAR(50);

ALTER TABLE
    users
ALTER COLUMN
    "password" TYPE VARCHAR(255);

ALTER TABLE
    projects
ALTER COLUMN
    "name" TYPE VARCHAR(50);

ALTER TABLE
    projects
ALTER COLUMN
    "description" TYPE VARCHAR(500);

ALTER TABLE
    todos
ALTER COLUMN
    "name" TYPE VARCHAR(50);

ALTER TABLE
    todos
ALTER COLUMN
    "description" TYPE VARCHAR(50);
