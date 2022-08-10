-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
)VALUES(
    $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers 
WHERE id = $1 LIMIT 1;

-- name: UpdateTransfer :one
UPDATE transfers 
SET amount = $1
WHERE id = $2
RETURNING *;

-- name: ListTransfers :many
SELECT * FROM transfers
LIMIT $1 
OFFSET $2;

-- name: DeleteTransfer :exec
DELETE FROM transfers 
WHERE id = $1;
