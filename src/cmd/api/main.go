package main

import (
	"go-blog/internal/configs"
	"go-blog/internal/db"
	"go-blog/internal/store"
	"log"
)

const version = "1.0.0"

func main() {
	cfg := config{
		addr:   configs.GetString("ADDR", ":3002"),
		apiURL: configs.GetString("EXTERNAL_URL", "localhost:3002"),
		db: dbConfig{
			addr:         configs.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: configs.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: configs.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  configs.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: configs.GetString("ENV", "development"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
