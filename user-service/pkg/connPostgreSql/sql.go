package connpostgresql

import (
	"user-service/config"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func ConnectToDB(cfg config.Config) (*sql.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	var connDb *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		connDb, err = sql.Open("postgres", psqlString)
		if err == nil {
			err = connDb.Ping()
			if err == nil {
				return connDb, nil 
			}
		}

		log.Printf("Ma'lumotlar bazasiga ulanishda xatolik, qayta urinmoqda... (%d/5)\n", i+1)
		time.Sleep(2 * time.Second) 
	}

	return nil, fmt.Errorf("ma'lumotlar bazasiga ulanish mumkin emas: %v", err)
}
