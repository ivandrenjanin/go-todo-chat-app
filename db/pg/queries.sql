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
