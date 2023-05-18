// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package tutorial

import (
	"context"
	"time"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payment (
  customer_id, staff_id, rental_id, amount, payment_date
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING payment_id, customer_id, staff_id, rental_id, amount, payment_date
`

type CreatePaymentParams struct {
	CustomerID  int32
	StaffID     int32
	RentalID    int32
	Amount      string
	PaymentDate time.Time
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, createPayment,
		arg.CustomerID,
		arg.StaffID,
		arg.RentalID,
		arg.Amount,
		arg.PaymentDate,
	)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.CustomerID,
		&i.StaffID,
		&i.RentalID,
		&i.Amount,
		&i.PaymentDate,
	)
	return i, err
}

const deletePayment = `-- name: DeletePayment :exec
DELETE FROM payment
WHERE payment_id = $1
`

func (q *Queries) DeletePayment(ctx context.Context, paymentID int32) error {
	_, err := q.db.ExecContext(ctx, deletePayment, paymentID)
	return err
}

const getPayment = `-- name: GetPayment :one
SELECT payment_id, customer_id, staff_id, rental_id, amount, payment_date FROM payment
WHERE payment_id = $1 LIMIT 1
`

func (q *Queries) GetPayment(ctx context.Context, paymentID int32) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPayment, paymentID)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.CustomerID,
		&i.StaffID,
		&i.RentalID,
		&i.Amount,
		&i.PaymentDate,
	)
	return i, err
}

const listPayment = `-- name: ListPayment :many
SELECT payment_id, customer_id, staff_id, rental_id, amount, payment_date FROM payment
ORDER BY payment_id
`

func (q *Queries) ListPayment(ctx context.Context) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, listPayment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.PaymentID,
			&i.CustomerID,
			&i.StaffID,
			&i.RentalID,
			&i.Amount,
			&i.PaymentDate,
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

const updatePayment = `-- name: UpdatePayment :one
UPDATE payment
  set customer_id = $2,
  staff_id = $3,
  rental_id = $4,
  amount = $5,
  payment_date = $6
WHERE payment_id = $1
RETURNING payment_id, customer_id, staff_id, rental_id, amount, payment_date
`

type UpdatePaymentParams struct {
	PaymentID   int32
	CustomerID  int32
	StaffID     int32
	RentalID    int32
	Amount      string
	PaymentDate time.Time
}

func (q *Queries) UpdatePayment(ctx context.Context, arg UpdatePaymentParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, updatePayment,
		arg.PaymentID,
		arg.CustomerID,
		arg.StaffID,
		arg.RentalID,
		arg.Amount,
		arg.PaymentDate,
	)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.CustomerID,
		&i.StaffID,
		&i.RentalID,
		&i.Amount,
		&i.PaymentDate,
	)
	return i, err
}
