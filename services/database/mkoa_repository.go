package database

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"library/services/entity"

	"github.com/jackc/pgx/v4/pgxpool"
)

const tableName = "mikoa"

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
LIST (with pagination, sort, filter)
*/
func (r *MkoaRepository) List(ctx context.Context, filter *entity.MkoaFilter) ([]*entity.Mkoa, int32, error) {

	filterQuery, args, totalCount := r.mkoaPageFilter(ctx, filter)
	query := r.getMkoaBaseQuery() + " WHERE deleted_at IS NULL " + filterQuery

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
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
			return nil, totalCount, err
		}

		list = append(list, m)
	}

	return list, totalCount, rows.Err()
}

func (r *MkoaRepository) getMkoaBaseQuery() string {
	return `SELECT id, name, code, status, created_at, updated_at,
		deleted_at, created_by, updated_by, deleted_by
		FROM ` + tableName
}

func (r *MkoaRepository) mkoaPageFilter(ctx context.Context, filter *entity.MkoaFilter) (string, []any, int32) {

	var values []any
	var filterQuery string
	index := 1

	if filter != nil {
		if filter.Name != "" {
			filterQuery += " AND name ILIKE $" + strconv.Itoa(index)
			index++
			values = append(values, "%"+filter.Name+"%")
		}
		if filter.Code != "" {
			filterQuery += " AND code ILIKE $" + strconv.Itoa(index)
			index++
			values = append(values, "%"+filter.Code+"%")
		}
		if filter.Status != "" {
			filterQuery += " AND status = $" + strconv.Itoa(index)
			index++
			values = append(values, filter.Status)
		}
	}

	totalCount := r.getTotalCount(ctx, filterQuery, values)

	orderByClause := ""
	if filter != nil && filter.SortBy != "" {
		sortOrder := "ASC"
		if strings.ToUpper(filter.SortOrder) == "DESC" {
			sortOrder = "DESC"
		}
		allowedColumns := map[string]bool{
			"name": true, "code": true, "status": true,
			"created_at": true, "updated_at": true,
		}
		if allowedColumns[filter.SortBy] {
			orderByClause = fmt.Sprintf(" ORDER BY %s %s", filter.SortBy, sortOrder)
		}
	}
	if orderByClause == "" {
		orderByClause = " ORDER BY id ASC"
	}

	page, pageSize := int32(1), int32(10)
	if filter != nil {
		if filter.Page > 0 {
			page = filter.Page
		}
		if filter.PageSize > 0 {
			pageSize = filter.PageSize
		}
	}
	offset := (page - 1) * pageSize
	filterQuery += orderByClause + " LIMIT $" + strconv.Itoa(index)
	values = append(values, pageSize)
	index++
	filterQuery += " OFFSET $" + strconv.Itoa(index)
	values = append(values, offset)

	return filterQuery, values, totalCount
}

func (r *MkoaRepository) getTotalCount(ctx context.Context, innerQuery string, args []any) int32 {
	var totalCount int32
	query := `SELECT count(*) FROM ` + tableName + ` WHERE deleted_at IS NULL ` + innerQuery
	err := r.db.QueryRow(ctx, query, args...).Scan(&totalCount)
	if err != nil {
		return 0
	}
	return totalCount
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

/*
HARD DELETE
*/
func (r *MkoaRepository) HardDelete(ctx context.Context, id int64) error {

	query := `DELETE FROM ` + tableName + ` WHERE id = $1`

	cmd, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("no record deleted")
	}

	return nil
}
