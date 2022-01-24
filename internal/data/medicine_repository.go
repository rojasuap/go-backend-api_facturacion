package data

import (
	"context"
	"time"

	"github.com/orlmonteverde/go-postgres-microblog/pkg/medicine"
)

// MedicineRepository manages the operations with the database that
// correspond to the medicine model.
type MedicineRepository struct {
	Data *Data
}

// GetAll returns all medicine.
func (pr *MedicineRepository) GetAll(ctx context.Context) ([]medicine.Medicine, error) {
	q := `
	SELECT id, name, price, location, created_at, updated_at
		FROM medicines;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var medicines []medicine.Medicine
	for rows.Next() {
		var p medicine.Medicine
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Location, &p.CreatedAt, &p.UpdatedAt)
		medicines = append(medicines, p)
	}

	return medicines, nil
}

// GetOne returns one post by id.
func (pr *MedicineRepository) GetOne(ctx context.Context, id uint) (medicine.Medicine, error) {
	q := `
	SELECT id, name, price, location, created_at, updated_at
		FROM medicines WHERE id = $1;
	`

	row := pr.Data.DB.QueryRowContext(ctx, q, id)

	var p medicine.Medicine
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Location, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return medicine.Medicine{}, err
	}

	return p, nil
}

// Create adds a new post.
func (pr *MedicineRepository) Create(ctx context.Context, p *medicine.Medicine) error {
	q := `
	INSERT INTO medicines (name, price, location, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.Name, p.Price, p.Location, time.Now(), time.Now())

	err = row.Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a post by id.
func (pr *MedicineRepository) Update(ctx context.Context, id uint, p medicine.Medicine) error {
	q := `
	UPDATE medicines set name=$1, price=$2, location=$3, updated_at=$4
		WHERE id=$5;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.Name, p.Price, p.Location, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a post by id.
func (pr *MedicineRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM posts WHERE id=$1;`

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
