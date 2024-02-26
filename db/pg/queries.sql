-- name: User :one
SELECT
    *
FROM
    users
WHERE
    id = $1
    AND deleted_at IS NULL
LIMIT
    1;

-- name: UserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = $1
    AND deleted_at IS NULL
LIMIT
    1;

-- name: InsertUser :one
INSERT INTO
    users (first_name, last_name, email, PASSWORD)
VALUES
    ($1, $2, $3, $4) RETURNING id;

-- name: DeleteUser :exec
UPDATE
    users
SET
    deleted_at = NOW()
WHERE
    id = $1;

-- name: ProjectsByUserId :many
SELECT
    sqlc.embed(projects),
    sqlc.embed(project_assignments)
FROM
    projects
    JOIN project_assignments ON projects.id = project_assignments.project_id
WHERE
    project_assignments.user_id = $1
    AND projects.deleted_at IS NULL;

-- name: ProjectById :one
SELECT
    *
FROM
    projects
WHERE
    id = $1
    AND projects.deleted_at IS NULL;

-- name: DeleteProject :exec
UPDATE
    projects
SET
    deleted_at = NOW()
WHERE
    id = $1;

-- name: InsertProject :one
INSERT INTO
    projects (name, description, owner_id)
VALUES
    ($1, $2, $3) RETURNING *;

-- name: InsertProjectAssignment :one
INSERT INTO
    project_assignments (project_id, user_id, project_owner_id)
VALUES
    ($1, $2, $3) RETURNING *;
