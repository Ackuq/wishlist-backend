-- name: GetAccount :one
SELECT * FROM account
    WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account
    ORDER BY name;

-- name: CreateAccount :one
INSERT INTO account (
    name, email
) VALUES (
    $1, $2
) RETURNING *;
