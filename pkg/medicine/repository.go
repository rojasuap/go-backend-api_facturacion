package medicine

import "context"

// Repository handle the CRUD operations with Medicines.
type Repository interface {
	GetAll(ctx context.Context) ([]Medicine, error)
	GetOne(ctx context.Context, id uint) (Medicine, error)
	Create(ctx context.Context, medicine *Medicine) error
	Update(ctx context.Context, id uint, medicine Medicine) error
	Delete(ctx context.Context, id uint) error
}
