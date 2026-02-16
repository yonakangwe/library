package repository

import (
	"context"
	"fmt"
	"library/package/util"
	"library/services/database"
	"library/services/entity"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UniversityConn struct {
	pgConn *pgxpool.Pool
}

// NewInstance creates a new UniversityConn with a database connection
// Consider using dependency injection to pass pgConn from application startup
// to avoid creating new connections on every service call
func NewInstance() *UniversityConn {
	conn, err := database.Connect()
	if util.IsError(err) {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil
	}
	return &UniversityConn{
		pgConn: conn,
	}
}

func (connection *UniversityConn) Create(entity *entity.University) (int32, error) {
	var universityID int32
	query := `INSERT INTO universities (name, abbreviation, email, website, established_year, is_active, created_by, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := connection.pgConn.QueryRow(
		context.Background(),
		query,
		entity.Name,
		entity.Abbreviation,
		entity.Email,
		entity.Website,
		entity.EstablishedYear,
		entity.IsActive,
		entity.CreatedBy,
		entity.CreatedAt,
	).Scan(&universityID)
	if util.IsError(err) {
		return 0, err
	}
	return universityID, nil
}
