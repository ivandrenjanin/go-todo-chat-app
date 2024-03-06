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
    public_id = $1
    AND projects.deleted_at IS NULL;

-- name: DeleteProject :exec
UPDATE
    projects
SET
    deleted_at = NOW()
WHERE
    public_id = $1;

-- name: UpdateProject :exec
UPDATE
    projects
SET
    name = $1,
    description = $2
WHERE
    owner_id = $3
    AND public_id = $4;

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

-- name: InsertProjectInvitation :one
INSERT INTO
    project_invitations (project_id, email, token, sent_at, expires_at)
VALUES
    ($1, $2, $3, $4, $5) RETURNING *;

-- name: InsertProjectTodoStates :exec
INSERT INTO
    project_todo_states (name, item_order, project_id)
VALUES
    ('ready', 0, $1),
    ('in-progress', 1, $1),
    ('done', 2, $1);
