package server

import (
	"github.com/Programming-Judge/Server/internal/database"
	"github.com/Programming-Judge/Server/internal/store"
)

func Start() {
	store.SetDBConnection(database.NewDBOptions())

	router := setRouter()
	// Start listening and serving requests
	router.Run(":8080")
}
