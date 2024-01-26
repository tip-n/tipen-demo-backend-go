package envconf

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type DatabaseEnv struct {
	DBUsername string `validate:"required"`
	DBPassword string `validate:"required"`
	DBName     string `validate:"required"`
	DBHostname string `validate:"required"`
	DBPort     string `validate:"required"`
}

type Env struct {
	Database DatabaseEnv
}

func ReadEnv() Env {
	if os.Getenv("MODE") != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("err loading: %v", err)
		}
	}

	e := Env{
		Database: DatabaseEnv{
			DBUsername: os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASS"),
			DBName:     os.Getenv("DB_NAME"),
			DBHostname: os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
		},
	}
	// if err := pkg.Validator.Struct(e); err != nil {
	// 	log.Fatalf("error: %v", err)
	// }

	return e
}

func (db DatabaseEnv) Url() string {
	// url := "postgresql://" +
	// 	db.DBUsername + ":" +
	// 	db.DBPassword + "@" +
	// 	db.DBHostname + ":" +
	// 	db.DBPort + "/" +
	// 	db.DBName +
	// 	"?sslmode=disable"

	url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		db.DBHostname,
		db.DBUsername,
		db.DBPassword,
		db.DBName,
		db.DBPort,
	)
	return url
}

func (db DatabaseEnv) Init() (*pgxpool.Pool, error) {
	var err error
	pool, err := pgxpool.New(context.Background(), db.Url())
	if err != nil {
		return nil, err
	}
	return pool, nil
}
