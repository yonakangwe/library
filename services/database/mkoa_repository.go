package database

import (
	"context"
	"errors"
	"time"

	"library/services/entity"

	"github.com/jackc/pgx/v4/pgxpool"
)

type MkoaRepository struct {
	db *pgxpool.Pool
}

func NewMkoaRepository(db *pgxpool.Pool) *MkoaRepository {
	return &MkoaRepository{db: db}
}

/*
INSERT
*/
func (r *MkoaRepository) Insert(ctx context.Context, m *entity.Mkoa) error {

	if err := m.ValidateCreate(); err != nil {
		return err
	}

	query := `
		INSERT INTO mikoa (name, code, status, created_at, created_by)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	return r.db.QueryRow(
		ctx,
		query,
		m.Name,
		m.Code,
		m.Status,
		m.CreatedAt,
		m.CreatedBy,
	).Scan(&m.ID)
}

/*
GET ALL (excluding soft deleted)
*/
func (r *MkoaRepository) GetAll(ctx context.Context) ([]*entity.Mkoa, error) {

	query := `
		SELECT id, name, code, status, created_at, updated_at,
		       deleted_at, created_by, updated_by, deleted_by
		FROM mikoa
		WHERE deleted_at IS NULL
		ORDER BY id;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*entity.Mkoa

	for rows.Next() {
		m := new(entity.Mkoa)

		if err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.Code,
			&m.Status,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.DeletedAt,
			&m.CreatedBy,
			&m.UpdatedBy,
			&m.DeletedBy,
		); err != nil {
			return nil, err
		}

		list = append(list, m)
	}

	return list, rows.Err()
}

/*
GET BY ID
*/
func (r *MkoaRepository) GetByID(ctx context.Context, id int64) (*entity.Mkoa, error) {

	query := `
		SELECT id, name, code, status, created_at, updated_at,
		       deleted_at, created_by, updated_by, deleted_by
		FROM mikoa
		WHERE id = $1 AND deleted_at IS NULL;
	`

	m := new(entity.Mkoa)

	err := r.db.QueryRow(ctx, query, id).Scan(
		&m.ID,
		&m.Name,
		&m.Code,
		&m.Status,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.DeletedAt,
		&m.CreatedBy,
		&m.UpdatedBy,
		&m.DeletedBy,
	)

	if err != nil {
		return nil, err
	}

	return m, nil
}

/*
UPDATE
*/
func (r *MkoaRepository) Update(ctx context.Context, m *entity.Mkoa) error {

	if err := m.ValidateUpdate(); err != nil {
		return err
	}

	now := time.Now()
	m.UpdatedAt = &now

	query := `
		UPDATE mikoa
		SET name = $1,
		    code = $2,
		    status = $3,
		    updated_at = $4,
		    updated_by = $5
		WHERE id = $6 AND deleted_at IS NULL;
	`

	cmd, err := r.db.Exec(
		ctx,
		query,
		m.Name,
		m.Code,
		m.Status,
		m.UpdatedAt,
		m.UpdatedBy,
		m.ID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("no record updated")
	}

	return nil
}

/*
SOFT DELETE
*/
func (r *MkoaRepository) SoftDelete(ctx context.Context, m *entity.Mkoa) error {

	if err := m.ValidateDelete(); err != nil {
		return err
	}

	now := time.Now()
	m.DeletedAt = &now

	query := `
		UPDATE mikoa
		SET deleted_at = $1,
		    deleted_by = $2
		WHERE id = $3 AND deleted_at IS NULL;
	`

	cmd, err := r.db.Exec(
		ctx,
		query,
		m.DeletedAt,
		m.DeletedBy,
		m.ID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("no record deleted")
	}

	return nil
}
