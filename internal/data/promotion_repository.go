package data

import (
	"context"
	"time"

	"github.com/rojasuap/go-backend-api_facturacion/pkg/promotion"
)

// PromotionsRepository manages the operations with the database that
// correspond to the post model.
type PromotionRepository struct {
	Data *Data
}

// GetAll returns all posts.
func (pr *PromotionRepository) GetAll(ctx context.Context) ([]promotion.Promotion, error) {
	q := `
	SELECT id, description, percentage, start_date, end_date, created_at, updated_at
		FROM promotions;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var promotions []promotion.Promotion
	for rows.Next() {
		var p promotion.Promotion
		rows.Scan(&p.ID, &p.Description, &p.Percentage, &p.Start_date, &p.End_date, &p.CreatedAt, &p.UpdatedAt)
		promotions = append(promotions, p)
	}

	return promotions, nil
}

// GetOne returns one promotion by id.
func (pr *PromotionRepository) GetOne(ctx context.Context, id uint) (promotion.Promotion, error) {
	q := `
	SELECT id, description, percentage, start_date, end_date, created_at, updated_at
		FROM promotions WHERE id = $1;
	`

	row := pr.Data.DB.QueryRowContext(ctx, q, id)

	var p promotion.Promotion
	err := row.Scan(&p.ID, &p.Description, &p.Percentage, &p.Start_date, &p.End_date, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return promotion.Promotion{}, err
	}

	return p, nil
}

// Create adds a new promotion.
func (pr *PromotionRepository) Create(ctx context.Context, p *promotion.Promotion) error {
	q := `
	INSERT INTO promotions (description, percentage, start_date, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.Description, p.Percentage, p.Start_date, p.End_date, time.Now(), time.Now())

	err = row.Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a promotion by id.
func (pr *PromotionRepository) Update(ctx context.Context, id uint, p promotion.Promotion) error {
	q := `
	UPDATE promotions set description=$1, percentage=$2, start_date=$3, end_date=$4 , updated_at=$5
		WHERE id=$6;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.Description, p.Percentage, p.Start_date, p.End_date, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a promotion by id.
func (pr *PromotionRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM promotions WHERE id=$1;`

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
