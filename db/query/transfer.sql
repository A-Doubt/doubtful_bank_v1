-- name: CreateTransfer :one
INSERT INTO transfer (
  from_account_id,
  to_account_id,
  amount,
  note
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfer
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfer
WHERE from_account_id = $1
OR to_account_id = $2
ORDER BY created_at
LIMIT $3
OFFSET $4;
