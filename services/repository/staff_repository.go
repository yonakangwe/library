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

func getStaaffQuery() string {
	return `SELECT 
					 id,
					 fullName,
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
