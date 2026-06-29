package main

import (
	"log"

	"github.com/aidostt/task-manager/internal/config"
	db2 "github.com/aidostt/task-manager/pkg/db"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db2.Connect(cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Connected to database, port:", cfg.ServerPort)
}
