package mkoa

import (
	"context"
	"library/services/entity"
)

type Repository interface {
	Insert(ctx context.Context, m *entity.Mkoa) error
	GetAll(ctx context.Context) ([]*entity.Mkoa, error)
	GetByID(ctx context.Context, id int64) (*entity.Mkoa, error)
	Update(ctx context.Context, m *entity.Mkoa) error
	SoftDelete(ctx context.Context, m *entity.Mkoa) error
}

type Usecase interface {
	Create(ctx context.Context, name, code string, createdBy int64) (*entity.Mkoa, error)
	GetAll(ctx context.Context) ([]*entity.Mkoa, error)
	GetByID(ctx context.Context, id int64) (*entity.Mkoa, error)
	Update(ctx context.Context, m *entity.Mkoa) error
	Delete(ctx context.Context, id int64, deletedBy int64) error
}
