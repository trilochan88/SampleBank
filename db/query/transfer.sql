-- name: CreateNewTransfer :one
INSERT INTO transfer(from_account_id,
                     to_account_id,
                     amount,
                     created_at)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetTransferById :one
select * from transfer where transfer_id = $1 FOR NO KEY UPDATE;

-- name: ListTransferByFromAccountId :many
select * from transfer where from_account_id = $1 limit $2 OFFSET $3 FOR NO KEY UPDATE;

-- name: ListTransferByToAccountId :many
select * from transfer where to_account_id = $1 limit $2 OFFSET $3 FOR NO KEY UPDATE;

-- name: DeleteTransferByTransferId :exec
Delete from transfer where transfer_id = $1;

-- name: DeleteTransferForAccountId :exec
Delete from transfer where from_account_id = $1 and to_account_id = $2;