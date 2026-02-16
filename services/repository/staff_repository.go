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
	"github.com/k0kubun/pp"
)

type StaffConn struct {
	conn *pgxpool.Pool
}

func NewStaff() *StaffConn {
	conn, err := database.Connect()
	if util.IsError(err) {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &StaffConn{
		conn: conn,
	}
}

var StaffTableName string = "staff"

func getStaffQuery() string {
	return `SELECT 
					 id,
					 full_name,
					 email,
					 phone,
					 username,
					 password_hash,
					 created_by, 
					 created_at, 
					 updated_by, 
					 updated_at, 
					 deleted_by, 
					 deleted_at 
				 FROM ` + StaffTableName
}

func (con *StaffConn) Create(e *entity.Staff) (int32, error) {
	pp.Println(e)
	var StaffID int32
	query := `INSERT INTO ` + StaffTableName + ` (full_name,email,phone,username,password_hash, created_by, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := con.conn.QueryRow(context.Background(), query, e.FullName, e.Email, e.Phone, e.Username, e.PasswordHash, e.CreatedBy, time.Now()).Scan(&StaffID)
	if util.IsError(err) {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return StaffID, nil
		}
		log.Errorf("error creating staff from table %v: %v", StaffTableName, err)
	}
	return StaffID, err
}

func (con *StaffConn) Update(e *entity.Staff) (int32, error) {
	query := `UPDATE ` + StaffTableName + ` SET 
										 full_name = $1,
										 email = $2,
										 phone = $3,
										 username = $4,
										 updated_by = $5, 
										 updated_at = $6 
										 WHERE id = $7`
	_, err := con.conn.Exec(context.Background(), query, e.FullName, e.Email, e.Phone, e.Username, e.UpdatedBy, time.Now(), e.ID)
	if util.IsError(err) {
		log.Errorf("error updating  from table %v by id: %v", StaffTableName, err)
		return 0, err
	}
	return e.ID, nil
}

func (con *StaffConn) Delete(e *entity.Staff) (int32, error) {
	query := `DELETE FROM ` + StaffTableName + ` WHERE id = $1`
	_, err := con.conn.Exec(context.Background(), query, e.ID)
	if util.IsError(err) {
		log.Errorf("error deleting from table %v by id: %v", StaffTableName, err)
		return 0, err
	}
	return e.ID, nil
}

func (con *StaffConn) Get(id int32) (*entity.Staff, error) {
	query := getStaffQuery() + " WHERE id = $1"
	row := con.conn.QueryRow(context.Background(), query, id)

	staff := &entity.Staff{}
	err := row.Scan(
		&staff.ID,
		&staff.FullName,
		&staff.Email,
		&staff.Phone,
		&staff.Username,
		&staff.PasswordHash,
		&staff.CreatedBy,
		&staff.CreatedAt,
		&staff.UpdatedBy,
		&staff.UpdatedAt,
		&staff.DeletedBy,
		&staff.DeletedAt,
	)

	if util.IsError(err) {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return nil, nil
		}
		log.Errorf("error getting staff from table %v by id: %v", StaffTableName, err)
		return nil, err
	}

	return staff, nil
}
