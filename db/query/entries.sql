-- name: CreateNewEntry :one
INSERT INTO entries(account_id,
                     amount,
                     created_at)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetEntryById :one
SELECT * FROM entries WHERE entries_id = $1 limit 1 FOR NO KEY UPDATE;

-- name: ListEntriesByAccountId :many
SELECT * FROM entries WHERE account_id = $1 limit $2 OFFSET $3 FOR NO KEY UPDATE;

-- name: DeleteEntry :exec
DELETE FROM entries WHERE entries_id = $1;
