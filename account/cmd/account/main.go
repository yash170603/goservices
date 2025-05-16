package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"github.com/yash170603/goservices/account"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

//look for an environment variable called DATABASE_URL, and put its value into this DatabaseURL field.”
//cfg.DatabaseURL = "postgres://user:pass@host:5432/db"
//Create a variable named cfg, and its type is Config.”

// So cfg is now a struct that looks like:
// Config{
//     DatabaseURL: "",
// }

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal("Error processing env config:", err)
	}

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println("Error connecting to database:", err)
		}
		return
	})

	defer r.Close()
	log.Println("Connected to database on port 8080")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}
