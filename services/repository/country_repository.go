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
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CountryConn struct {
	conn *pgxpool.Pool
}

func NewCountry() *CountryConn {
	conn, err := database.Connect()
	if util.IsError(err) {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &CountryConn{
		conn: conn,
	}
}

var countryTableName string = "countries"

func getCountryQuery() string {
	return `SELECT 
					 id,
					 name, 
					 iso_code,
					 phone_code,
					 is_active,
					 created_by, 
					 created_at, 
					 updated_by, 
					 updated_at, 
					 deleted_by, 
					 deleted_at 
				 FROM ` + countryTableName
}

func (con *CountryConn) Create(e *entity.Country) (int32, error) {
	var CountryID int32
	query := `INSERT INTO ` + countryTableName + ` (name, iso_code, phone_code, is_active, created_by, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := con.conn.QueryRow(context.Background(), query, e.Name, e.IsoCode, e.PhoneCode, e.IsActive, e.CreatedBy, time.Now()).Scan(&CountryID)
	if util.IsError(err) {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return CountryID, nil
		}
		log.Errorf("error creating country from table %v: %v", countryTableName, err)
	}
	return CountryID, err
}

func (con *CountryConn) List(filter *entity.CountryFilter) ([]*entity.Country, int32, error) {
	var id pgtype.Int4
	var name, isoCode pgtype.GenericText
	var phoneCode pgtype.Int2
	var isActive pgtype.Bool
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var createdAt, updatedAt, deletedAt pgtype.Timestamp
	var data []*entity.Country

	joinQuery, args, totalCount := con.CountryPageFilter(filter)
	query := getCountryQuery() + " WHERE deleted_at IS NULL " + joinQuery
	rows, err := con.conn.Query(context.Background(), query, args...)
	if util.IsError(err) {
		log.Errorf("error querying  table %v: %v", countryTableName, err)
		return nil, totalCount, err
	}
	for rows.Next() {
		if err := rows.Scan(&id, &name, &isoCode, &phoneCode, &isActive, &createdBy, &createdAt, &updatedBy, &updatedAt, &deletedBy, &deletedAt); util.IsError(err) {
			log.Errorf("error scanning from table %v : %v", countryTableName, err)
			return nil, totalCount, err
		}
		CountryData := &entity.Country{
			ID:        id.Int,
			Name:      name.String,
			IsoCode:   isoCode.String,
			PhoneCode: phoneCode.Int,
			IsActive:  isActive.Bool,
			CreatedBy: createdBy.Int,
			CreatedAt: createdAt.Time,
			UpdatedBy: updatedBy.Int,
			UpdatedAt: updatedAt.Time,
			DeletedBy: deletedBy.Int,
			DeletedAt: deletedAt.Time,
		}
		data = append(data, CountryData)
	}
	return data, totalCount, err
}

func (con *CountryConn) Get(id int32) (*entity.Country, error) {
	var name, isoCode pgtype.GenericText
	var phoneCode pgtype.Int2
	var isActive pgtype.Bool
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var createdAt, updatedAt, deletedAt pgtype.Timestamp
	query := getCountryQuery() + ` WHERE deleted_at IS NULL AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(&id, &name, &isoCode, &phoneCode, &isActive, &createdBy, &createdAt, &updatedBy, &updatedAt, &deletedBy, &deletedAt)
	if util.IsError(err) {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return nil, nil
		}
		log.Errorf("error getting  table %v: %v", countryTableName, err)
		return nil, err
	}
	Country := &entity.Country{
		ID:        id,
		Name:      name.String,
		IsoCode:   isoCode.String,
		PhoneCode: phoneCode.Int,
		IsActive:  isActive.Bool,
		CreatedBy: createdBy.Int,
		CreatedAt: createdAt.Time,
		UpdatedBy: updatedBy.Int,
		UpdatedAt: updatedAt.Time,
		DeletedBy: deletedBy.Int,
		DeletedAt: deletedAt.Time,
	}
	return Country, err
}

func (con *CountryConn) Update(e *entity.Country) error {
	query := `UPDATE ` + countryTableName + ` SET 
										 name = $1, 
										 iso_code = $2, 
										 phone_code = $3, 
										 is_active = $4, 
										 updated_by = $5, 
										 updated_at = $6 
										 WHERE id = $7`
	_, err := con.conn.Exec(context.Background(), query, e.Name, e.IsoCode, e.PhoneCode, e.IsActive, e.UpdatedBy, time.Now(), e.ID)
	if util.IsError(err) {
		log.Errorf("error updating  from table %v by id: %v", countryTableName, err)
	}
	return err
}

func (con *CountryConn) SoftDelete(id, deletedBy int32) error {
	query := `UPDATE ` + countryTableName + ` SET 
									 deleted_by = $1, 
									 deleted_at = $2
									 WHERE id = $3`
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now(), id)
	if util.IsError(err) {
		log.Errorf("error soft delete  from table %v by id: %v", countryTableName, err)
	}
	return err
}

func (con *CountryConn) HardDelete(id int32) error {
	query := `DELETE FROM ` + countryTableName + ` WHERE id = $1`
	_, err := con.conn.Exec(context.Background(), query, id)
	if util.IsError(err) {
		log.Errorf("error hard delete from table %v by id: %v", countryTableName, err)
	}
	return err
}

func (con *CountryConn) CountryPageFilter(filter *entity.CountryFilter) (string, []any, int32) {
	var values []any
	var filterQuery string
	index := 1

	if filter.Name != "" {
		filterQuery += " AND name ILIKE $" + strconv.Itoa(index)
		index++
		values = append(values, "%"+filter.Name+"%")
	}

	totalCount := con.GetCountryTotalCount(filterQuery, values)
	orderByClause := ""
	if filter.SortBy != "" {
		sortOrder := "ASC"
		if strings.ToUpper(filter.SortOrder) == "DESC" {
			sortOrder = "DESC"
		}
		allowedColumns := map[string]bool{
			"name":       true,
			"iso_code":   true,
			"phone_code": true,
			"is_active":  true,
			"created_at": true,
			"updated_at": true,
		}
		if allowedColumns[filter.SortBy] {
			orderByClause = fmt.Sprintf(" ORDER BY %s %s", filter.SortBy, sortOrder)
		}
	}

	paginationQuery, args := GetPaginationQuery(filter.Page, filter.PageSize, index, values)
	finalQuery := filterQuery + orderByClause + paginationQuery

	return finalQuery, args, totalCount
}

func (con *CountryConn) GetCountryTotalCount(innerQuery string, args []any) int32 {
	var totalCount int32
	query := ` SELECT count(*) FROM ` + countryTableName + ` WHERE deleted_at IS NULL ` + innerQuery
	err := con.conn.QueryRow(context.Background(), query, args...).Scan(&totalCount)
	if err != nil {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return 0
		}
		log.Errorf("error getting total count from %v : %v", countryTableName, err)
	}
	return totalCount
}
