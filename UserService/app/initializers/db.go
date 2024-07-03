package initializers

import (
	"GwentMicroservices/UserService/app/helpers"
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

var (
	DB_PORT     = helpers.GetEnv("DB_PORT")
	DB_HOST     = helpers.GetEnv("DB_HOST")
	DB_NAME     = helpers.GetEnv("DB_NAME")
	DB_USER     = helpers.GetEnv("DB_USER")
	DB_PASSWORD = helpers.GetEnv("DB_PASSWORD")
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
	//runMigrations()
	log.Info("DB connected successfully")
	return pool
}

func runMigrations() {
	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	cmd := exec.Command("migrate", "-path", "db/migrations", "-database", psqlInfo, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Panicf("Migrations error: %v", err)
	}

	log.Info("Migrations have passed!")
}
