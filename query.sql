-- name: GetPayment :one
SELECT * FROM payment
WHERE payment_id = $1 LIMIT 1;

-- name: ListPayment :many
SELECT * FROM payment
ORDER BY payment_id;

-- name: CreatePayment :one
INSERT INTO payment (
  customer_id, staff_id, rental_id, amount, payment_date
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeletePayment :exec
DELETE FROM payment
WHERE payment_id = $1;

-- name: UpdatePayment :one
UPDATE payment
  set customer_id = $2,
  staff_id = $3,
  rental_id = $4,
  amount = $5,
  payment_date = $6
WHERE payment_id = $1
RETURNING *;
