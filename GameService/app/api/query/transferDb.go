package query

import "github.com/jackc/pgx/v5/pgxpool"

var DB *pgxpool.Pool

func TransferDB(db *pgxpool.Pool) {
	DB = db
}
