// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: entries.sql

package db

import (
	"context"
	"time"
)

const createNewEntry = `-- name: CreateNewEntry :one
INSERT INTO entries(account_id,
                     amount,
                     created_at)
VALUES ($1, $2, $3) RETURNING entries_id, account_id, amount, created_at
`

type CreateNewEntryParams struct {
	AccountID int64     `json:"account_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateNewEntry(ctx context.Context, arg CreateNewEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createNewEntry, arg.AccountID, arg.Amount, arg.CreatedAt)
	var i Entry
	err := row.Scan(
		&i.EntriesID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM entries WHERE entries_id = $1
`

func (q *Queries) DeleteEntry(ctx context.Context, entriesID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntry, entriesID)
	return err
}

const getEntryById = `-- name: GetEntryById :one
SELECT entries_id, account_id, amount, created_at FROM entries WHERE entries_id = $1 limit 1 FOR NO KEY UPDATE
`

func (q *Queries) GetEntryById(ctx context.Context, entriesID int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntryById, entriesID)
	var i Entry
	err := row.Scan(
		&i.EntriesID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntriesByAccountId = `-- name: ListEntriesByAccountId :many
SELECT entries_id, account_id, amount, created_at FROM entries WHERE account_id = $1 limit $2 OFFSET $3 FOR NO KEY UPDATE
`

type ListEntriesByAccountIdParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) ListEntriesByAccountId(ctx context.Context, arg ListEntriesByAccountIdParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntriesByAccountId, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.EntriesID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
