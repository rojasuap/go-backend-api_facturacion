package promotion

import "context"

// Repository handle the CRUD operations with Promotions.
type Repository interface {
	GetAll(ctx context.Context) ([]Promotion, error)
	GetOne(ctx context.Context, id uint) (Promotion, error)
	//GetByUser(ctx context.Context, userID uint) ([]Post, error)
	Create(ctx context.Context, promotion *Promotion) error
	Update(ctx context.Context, id uint, promotion Promotion) error
	Delete(ctx context.Context, id uint) error
}
