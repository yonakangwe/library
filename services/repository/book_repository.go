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

type BookConn struct {
	conn *pgxpool.Pool
}

func NewBook() *BookConn {
	conn, err := database.Connect()
	if util.IsError(err) {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &BookConn{
		conn: conn,
	}
}

var tableName string = "books"

func getBookQuery() string {
	return `SELECT 
					 id,
					 title, 
					 author,
					 isbn,
					 status,
					 created_by, 
					 created_at, 
					 updated_by, 
					 updated_at, 
					 deleted_by, 
					 deleted_at 
				 FROM ` + tableName
}

func (con *BookConn) Create(e *entity.Book) (int32, error) {
	var BookID int32
	query := `INSERT INTO ` + tableName + ` (title, author, isbn, status, created_by, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := con.conn.QueryRow(context.Background(), query, e.Title, e.Author, e.Isbn, e.Status, e.CreatedBy, time.Now()).Scan(&BookID)
	if util.IsError(err) {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return BookID, nil
		}
		log.Errorf("error creating book from table %v: %v", tableName, err)
	}
	return BookID, err
}
