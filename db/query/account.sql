-- name: CreateAccount :one
INSERT INTO accounts(owner,
                     balance,
                     currency,
                     created_at)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts where account_id = $1 limit 1 FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts order by account_id LIMIT  $1 OFFSET $2 FOR NO KEY UPDATE;

-- name: UpdateAccount :one
UPDATE accounts SET balance = $2 WHERE account_id = $1 RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE account_id = sqlc.arg(account_id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE account_id = $1;