package repository

import (
	"context"
	"fmt"
	"library/package/log"
	"library/package/util"
	"library/services/database"
	"library/services/entity"
	"library/services/error_message"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type RoleConn struct {
	conn *pgxpool.Pool
}

func NewRole() *RoleConn {
	conn, err := database.Connect()
	if util.IsError(err) {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &RoleConn{
		conn: conn,
	}
}

var tableName string = "roles"

func getRoleQuery() string {
	return `SELECT 
					 id,
					 name, 
					 description,
					 created_by, 
					 created_at, 
					 updated_by, 
					 updated_at, 
					 deleted_by, 
					 deleted_at 
				 FROM ` + tableName
}

func (con *RoleConn) Create(e *entity.Role) (int32, error) {
	var RoleID int32
	query := `INSERT INTO ` + tableName + ` (name, description, created_by, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := con.conn.QueryRow(context.Background(), query, e.Name, e.Description, e.CreatedBy, time.Now()).Scan(&RoleID)
	if util.IsError(err) {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return RoleID, nil
		}
		log.Errorf("error creating role from table %v: %v", tableName, err)
	}
	return RoleID, err
}

// func (con *RoleConn) List(filter *entity.RoleFilter) ([]*entity.Role, int32, error) {
// 	var id pgtype.Int4
// 	var name, description pgtype.GenericText
// 	var createdBy, updatedBy, deletedBy pgtype.Int4
// 	var createdAt, updatedAt, deletedAt pgtype.Timestamp
// 	var data []*entity.Role

// 	joinQuery, args, totalCount := con.RolePageFilter(filter)
// 	query := getRoleQuery() + " WHERE deleted_at IS NULL " + joinQuery
// 	rows, err := con.conn.Query(context.Background(), query, args...)
// 	if util.IsError(err) {
// 		log.Errorf("error querying  table %v: %v", tableName, err)
// 		return nil, totalCount, err
// 	}
// 	for rows.Next() {
// 		if err := rows.Scan(&id, &name, &description, &createdBy, &createdAt, &updatedBy, &updatedAt, &deletedBy, &deletedAt); util.IsError(err) {
// 			log.Errorf("error scanning from table %v : %v", tableName, err)
// 			return nil, totalCount, err
// 		}
// 		RoleData := &entity.Role{
// 			ID:          id.Int,
// 			Name:        name.String,
// 			Description: description.String,
// 			CreatedBy:   createdBy.Int,
// 			CreatedAt:   createdAt.Time,
// 			UpdatedBy:   updatedBy.Int,
// 			UpdatedAt:   updatedAt.Time,
// 			DeletedBy:   deletedBy.Int,
// 			DeletedAt:   deletedAt.Time,
// 		}
// 		data = append(data, RoleData)
// 	}
// 	return data, totalCount, err
// }

// func (con *RoleConn) Get(id int32) (*entity.Role, error) {
// 	var name, description pgtype.GenericText
// 	var createdBy, updatedBy, deletedBy pgtype.Int4
// 	var createdAt, updatedAt, deletedAt pgtype.Timestamp
// 	query := getRoleQuery() + ` WHERE deleted_at IS NULL AND id = $1`
// 	err := con.conn.QueryRow(context.Background(), query, id).Scan(&id, &name, &description, &createdBy, &createdAt, &updatedBy, &updatedAt, &deletedBy, &deletedAt)
// 	if util.IsError(err) {
// 		if err.Error() == error_message.ErrNoResultSet.Error() {
// 			return nil, nil
// 		}
// 		log.Errorf("error getting  table %v: %v", tableName, err)
// 		return nil, err
// 	}
// 	Role := &entity.Role{
// 		ID:          id,
// 		Name:        name.String,
// 		Description: description.String,
// 		CreatedBy:   createdBy.Int,
// 		CreatedAt:   createdAt.Time,
// 		UpdatedBy:   updatedBy.Int,
// 		UpdatedAt:   updatedAt.Time,
// 		DeletedBy:   deletedBy.Int,
// 		DeletedAt:   deletedAt.Time,
// 	}
// 	return Role, err
// }

// func (con *RoleConn) Update(e *entity.Role) error {
// 	query := `UPDATE ` + tableName + ` SET
// 										 name = $1,
// 										 description = $2,
// 										 updated_by = $3,
// 										 updated_at = $4
// 										 WHERE id = $5`
// 	_, err := con.conn.Exec(context.Background(), query, e.Name, e.Description, e.UpdatedBy, time.Now(), e.ID)
// 	if util.IsError(err) {
// 		log.Errorf("error updating  from table %v by id: %v", tableName, err)
// 	}
// 	return err
// }

// func (con *RoleConn) SoftDelete(id, deletedBy int32) error {
// 	query := `UPDATE ` + tableName + ` SET
// 									 deleted_by = $1,
// 									 deleted_at = $2
// 									 WHERE id = $3`
// 	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now(), id)
// 	if util.IsError(err) {
// 		log.Errorf("error soft delete  from table %v by id: %v", tableName, err)
// 	}
// 	return err
// }

// func (con *RoleConn) HardDelete(id int32) error {
// 	query := `DELETE FROM ` + tableName + ` WHERE id = $1`
// 	_, err := con.conn.Exec(context.Background(), query, id)
// 	if util.IsError(err) {
// 		log.Errorf("error hard delete from table %v by id: %v", tableName, err)
// 	}
// 	return err
// }

// func (con *RoleConn) RolePageFilter(filter *entity.RoleFilter) (string, []any, int32) {
// 	var values []any
// 	var filterQuery string
// 	index := 1

// 	if filter.Name != "" {
// 		filterQuery += " AND name ILIKE $" + strconv.Itoa(index)
// 		index++
// 		values = append(values, "%"+filter.Name+"%")
// 	}

// 	totalCount := con.GetTotalCount(filterQuery, values)
// 	orderByClause := ""
// 	if filter.SortBy != "" {
// 		sortOrder := "ASC"
// 		if strings.ToUpper(filter.SortOrder) == "DESC" {
// 			sortOrder = "DESC"
// 		}
// 		allowedColumns := map[string]bool{
// 			"name":        true,
// 			"description": true,
// 			"created_at":  true,
// 			"updated_at":  true,
// 		}
// 		if allowedColumns[filter.SortBy] {
// 			orderByClause = fmt.Sprintf(" ORDER BY %s %s", filter.SortBy, sortOrder)
// 		}
// 	}

// 	paginationQuery, args := GetPaginationQuery(filter.Page, filter.PageSize, index, values)
// 	finalQuery := filterQuery + orderByClause + paginationQuery

// 	return finalQuery, args, totalCount
// }

// func (con *RoleConn) GetTotalCount(innerQuery string, args []any) int32 {
// 	var totalCount int32
// 	query := ` SELECT count(*) FROM ` + tableName + ` WHERE deleted_at IS NULL ` + innerQuery
// 	err := con.conn.QueryRow(context.Background(), query, args...).Scan(&totalCount)
// 	if err != nil {
// 		if err.Error() == error_message.ErrNoResultSet.Error() {
// 			return 0
// 		}
// 		log.Errorf("error getting total count from %v : %v", tableName, err)
// 	}
// 	return totalCount
// }
