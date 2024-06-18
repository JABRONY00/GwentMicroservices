package helpers

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ValidateName(name string, db *pgxpool.Pool) bool {
	res := db.QueryRow(context.Background(), "SELECT DISTINCT players_name FROM players WHERE players_name = $1", name)
	err := res.Scan()
	return err.Error() == "no rows in result set"
}
func ValidatePassword(password string) bool {
	return len([]rune(password)) >= 4
}
