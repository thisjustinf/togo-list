package cmd

import (
	"log"

	DB "github.com/thisjustinf/togo-list/internal/db"
	"github.com/thisjustinf/togo-list/pkg/server"
)

func main() {
	db, err := DB.InitDB()
	if err != nil {
		log.Panic(err)
	}

	server.Run(db)
}
