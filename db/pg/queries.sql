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
    project_assignments.user_id = $1;

-- name: ProjectById :one
SELECT
    *
FROM
    PROJECTS
WHERE
    id = $1;
