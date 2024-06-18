package initializers

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

var (
	DB_PORT     = "5432"
	DB_HOST     = "localhost"
	DB_NAME     = "gwent_db"
	DB_USER     = "GwentUser"
	DB_PASSWORD = "gwent"
)

func DbConnection() *pgxpool.Pool {
	postgreURL := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	pool, err := pgxpool.New(context.Background(), postgreURL)
	if err != nil {
		log.Panicf("DB Connection failed: %v", err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		log.Panicf("DB Ping failed: %v", err)
	}
	log.Info("DB connected successfully")
	return pool
}
