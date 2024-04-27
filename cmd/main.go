package main

import (
	"humoshop/api/server"
	"humoshop/pkg/db"
)

func main() {
	db.DatabaseConnect()
	defer db.DatabaseClose()

	server.StartRouter()
}
