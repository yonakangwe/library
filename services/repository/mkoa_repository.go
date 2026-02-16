package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"library/package/log"
	"library/package/util"
	"library/services/database"
	"library/services/entity"
	"library/services/error_message"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type MkoaConn struct {
	conn *pgxpool.Pool
}

func NewMkoa() *MkoaConn {
	conn, err := database.Connect()
	if util.IsError(err) {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &MkoaConn{
		conn: conn,
	}
}

var mkoaTableName string = "mikoa"

// ErrMkoaNotFound is returned when a mkoa record is not found (e.g. Get by id).
var ErrMkoaNotFound = errors.New("mkoa not found")

// ErrDBUnavailable is returned when the database connection is nil (e.g. PostgreSQL not running).
var ErrDBUnavailable = errors.New("database unavailable")

// ErrMkoaCodeExists is returned when creating or updating with a code that already exists (case-insensitive).
var ErrMkoaCodeExists = errors.New("mkoa code already exists")

func (con *MkoaConn) checkConn() error {
	if con == nil || con.conn == nil {
		return ErrDBUnavailable
	}
	return nil
}

func getMkoaQuery() string {
	return `SELECT id, name, code, status, created_at, updated_at,
		deleted_at, created_by, updated_by, deleted_by
		FROM ` + mkoaTableName
}

/*
GetByCode returns a mkoa by code (case-insensitive), only non-deleted. Used for uniqueness check.
*/
func (con *MkoaConn) GetByCode(code string) (*entity.Mkoa, error) {
	if err := con.checkConn(); err != nil {
		return nil, err
	}
	if code == "" {
		return nil, ErrMkoaNotFound
	}
	query := getMkoaQuery() + ` WHERE deleted_at IS NULL AND LOWER(code) = LOWER($1)`
	m := new(entity.Mkoa)
	err := con.conn.QueryRow(context.Background(), query, code).Scan(
		&m.ID, &m.Name, &m.Code, &m.Status, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt,
		&m.CreatedBy, &m.UpdatedBy, &m.DeletedBy,
	)
	if util.IsError(err) {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return nil, ErrMkoaNotFound
		}
		log.Errorf("error getting mkoa by code from table %v: %v", mkoaTableName, err)
		return nil, err
	}
	return m, nil
}

/*
CREATE
*/
func (con *MkoaConn) Create(e *entity.Mkoa) (int32, error) {
	if err := con.checkConn(); err != nil {
		return 0, err
	}
	var mkoaID int32
	if err := e.ValidateCreate(); err != nil {
		return 0, err
	}
	existing, err := con.GetByCode(e.Code)
	if err != nil && !errors.Is(err, ErrMkoaNotFound) {
		return 0, err
	}
	if existing != nil {
		return 0, ErrMkoaCodeExists
	}
	query := `INSERT INTO ` + mkoaTableName + ` (name, code, status, created_at, created_by) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = con.conn.QueryRow(context.Background(), query, e.Name, e.Code, e.Status, e.CreatedAt, entity.Int64PtrVal(e.CreatedBy)).Scan(&mkoaID)
	if util.IsError(err) {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return mkoaID, nil
		}
		log.Errorf("error creating mkoa from table %v: %v", mkoaTableName, err)
		return 0, err
	}
	e.ID = int64(mkoaID)
	return mkoaID, nil
}

/*
GET
*/
func (con *MkoaConn) Get(id int32) (*entity.Mkoa, error) {
	if err := con.checkConn(); err != nil {
		return nil, err
	}
	query := getMkoaQuery() + ` WHERE deleted_at IS NULL AND id = $1`
	m := new(entity.Mkoa)
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
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
	if util.IsError(err) {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return nil, ErrMkoaNotFound
		}
		log.Errorf("error getting mkoa from table %v: %v", mkoaTableName, err)
		return nil, err
	}
	return m, nil
}

/*
LIST (with pagination, sort, filter)
*/
func (con *MkoaConn) List(filter *entity.MkoaFilter) ([]*entity.Mkoa, int32, error) {
	if err := con.checkConn(); err != nil {
		return nil, 0, err
	}
	joinQuery, args, totalCount := con.mkoaPageFilter(filter)
	query := getMkoaQuery() + " WHERE deleted_at IS NULL " + joinQuery
	rows, err := con.conn.Query(context.Background(), query, args...)
	if util.IsError(err) {
		log.Errorf("error querying table %v: %v", mkoaTableName, err)
		return nil, totalCount, err
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
			log.Errorf("error scanning from table %v: %v", mkoaTableName, err)
			return nil, totalCount, err
		}
		list = append(list, m)
	}
	return list, totalCount, rows.Err()
}

func (con *MkoaConn) mkoaPageFilter(filter *entity.MkoaFilter) (string, []any, int32) {
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

	totalCount := con.getTotalCount(filterQuery, values)

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
	paginationQuery, args := GetPaginationQuery(page, pageSize, index, values)
	finalQuery := filterQuery + orderByClause + paginationQuery

	return finalQuery, args, totalCount
}

func (con *MkoaConn) getTotalCount(innerQuery string, args []any) int32 {
	if con.checkConn() != nil {
		return 0
	}
	var totalCount int32
	query := `SELECT count(*) FROM ` + mkoaTableName + ` WHERE deleted_at IS NULL ` + innerQuery
	err := con.conn.QueryRow(context.Background(), query, args...).Scan(&totalCount)
	if err != nil {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return 0
		}
		log.Errorf("error getting total count from %v: %v", mkoaTableName, err)
	}
	return totalCount
}

/*
UPDATE
*/
func (con *MkoaConn) Update(e *entity.Mkoa) error {
	if err := con.checkConn(); err != nil {
		return err
	}
	if err := e.ValidateUpdate(); err != nil {
		return err
	}
	existing, err := con.GetByCode(e.Code)
	if err != nil && !errors.Is(err, ErrMkoaNotFound) {
		return err
	}
	if existing != nil && existing.ID != e.ID {
		return ErrMkoaCodeExists
	}
	query := `UPDATE ` + mkoaTableName + ` SET
		name = $1,
		code = $2,
		status = $3,
		updated_at = $4,
		updated_by = $5
		WHERE id = $6 AND deleted_at IS NULL`
	now := time.Now()
	cmd, err := con.conn.Exec(context.Background(), query, e.Name, e.Code, e.Status, now, entity.Int64PtrVal(e.UpdatedBy), e.ID)
	if util.IsError(err) {
		log.Errorf("error updating from table %v: %v", mkoaTableName, err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		return ErrMkoaNotFound
	}
	return nil
}

/*
SOFT DELETE
*/
func (con *MkoaConn) SoftDelete(id, deletedBy int32) error {
	if err := con.checkConn(); err != nil {
		return err
	}
	m, err := con.Get(id)
	if err != nil {
		return err
	}
	db := int64(deletedBy)
	m.DeletedBy = &db
	if err := m.ValidateDelete(); err != nil {
		return err
	}
	query := `UPDATE ` + mkoaTableName + ` SET deleted_at = $1, deleted_by = $2 WHERE id = $3 AND deleted_at IS NULL`
	now := time.Now()
	cmd, err := con.conn.Exec(context.Background(), query, now, deletedBy, id)
	if util.IsError(err) {
		log.Errorf("error soft delete from table %v: %v", mkoaTableName, err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		return ErrMkoaNotFound
	}
	return nil
}

/*
HARD DELETE
*/
func (con *MkoaConn) HardDelete(id int32) error {
	if err := con.checkConn(); err != nil {
		return err
	}
	query := `DELETE FROM ` + mkoaTableName + ` WHERE id = $1`
	cmd, err := con.conn.Exec(context.Background(), query, id)
	if util.IsError(err) {
		log.Errorf("error hard delete from table %v: %v", mkoaTableName, err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		return ErrMkoaNotFound
	}
	return nil
}
