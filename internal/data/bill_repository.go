package data

import (
	"context"
	"time"

	"github.com/orlmonteverde/go-postgres-microblog/pkg/bill"
)

// PromotionsRepository manages the operations with the database that
// correspond to the post model.
type BillRepository struct {
	Data *Data
}

// GetAll returns all posts.
func (pr *BillRepository) GetAll(ctx context.Context) ([]bill.Bill, error) {
	q := `
	SELECT id, created_at, full_payment, updated_at
		FROM bills;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bills []bill.Bill
	for rows.Next() {
		var p bill.Bill
		rows.Scan(&p.ID, &p.CreatedAt, &p.Full_payment, &p.UpdatedAt)
		bills = append(bills, p)
	}

	return bills, nil
}

// GetOne returns one post by id.
func (pr *BillRepository) GetOne(ctx context.Context, id uint) (bill.Bill, error) {
	q := `
	SELECT id, created_at, full_payment, updated_at
		FROM bills WHERE id = $1;
	`

	row := pr.Data.DB.QueryRowContext(ctx, q, id)

	var p bill.Bill
	err := row.Scan(&p.ID, &p.CreatedAt, &p.Full_payment, &p.UpdatedAt)
	if err != nil {
		return bill.Bill{}, err
	}

	return p, nil
}

// GetAll returns all posts.
func (pr *BillRepository) GetAllDays(ctx context.Context) ([]bill.Bill, error) {
	q := `
	SELECT created_at, SUM(full_payment) as full_payment
		FROM bills
	 GROUP BY created_at 
	 ORDER BY created_at ASC;
	`
	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bills []bill.Bill
	for rows.Next() {
		var p bill.Bill
		rows.Scan(&p.CreatedAt, &p.Full_payment)
		bills = append(bills, p)
	}

	return bills, nil
}

// GetAll returns all posts.
func (pr *BillRepository) GetAllPayments(ctx context.Context) ([]bill.Bill, error) {
	q := `
	SELECT full_payment
		FROM bills
	 ORDER BY full_payment;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bills []bill.Bill
	for rows.Next() {
		var p bill.Bill
		rows.Scan(&p.Full_payment, &p.Full_payment)
		bills = append(bills, p)
	}

	return bills, nil
}

// Create adds a new post.
func (pr *BillRepository) Create(ctx context.Context, p *bill.Bill) error {
	q := `
	INSERT INTO bills ( created_at, full_payment, updated_at)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.CreatedAt, p.Full_payment, time.Now())

	err = row.Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a post by id.
func (pr *BillRepository) Update(ctx context.Context, id uint, p bill.Bill) error {
	q := `
	UPDATE bills set created_at=$1, full_payment=$2, updated_at=$3
		WHERE id=$4;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.CreatedAt, p.Full_payment, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a post by id.
func (pr *BillRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM bills WHERE id=$1;`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
