package bill

import "context"

// Repository handle the CRUD operations with Promotions.
type Repository interface {
	GetAll(ctx context.Context) ([]Bill, error)
	GetOne(ctx context.Context, id uint) (Bill, error)
	GetAllDays(ctx context.Context) ([]Bill, error)
	GetAllPayments(ctx context.Context) ([]Bill, error)
	Create(ctx context.Context, promotion *Bill) error
	Update(ctx context.Context, id uint, promotion Bill) error
	Delete(ctx context.Context, id uint) error
}
