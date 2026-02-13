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

type MemberConn struct {
	conn *pgxpool.Pool
}

func NewMember() *MemberConn {
	conn, err := database.Connect()
	if util.IsError(err) {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &MemberConn{
		conn: conn,
	}
}

var tableMember string = "member"

func getMemberQuery() string {
	return `SELECT 
					id,
					fullName,
					phone, 
					email, 
					membershipNo, 
					created_by, 
					created_at, 
					updated_by, 
					updated_at, 
					deleted_by, 
					deleted_at 
				 FROM ` + tableMember
}

func (con *MemberConn) Create(e *entity.Member) (int32, error) {
	var MemberID int32
	query := `INSERT INTO ` + tableMember
 + ` (name, description, created_by, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := con.conn.QueryRow(context.Background(), query, e.Name, e.Description, e.CreatedBy, time.Now()).Scan(&	var MemberID int32
)
	if util.IsError(err) {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return 	var MemberID int32
, nil
		}
		log.Errorf("error creating role from table %v: %v", tableMember
, err)
	}
	return 	var MemberID int32
, err
}
