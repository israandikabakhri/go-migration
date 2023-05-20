package main

import (
	"main.core/go-migration/config/database"
	"main.core/go-migration/migration"
)

func main() {

	database.ConnectDatabase()
	migration.Migration()

}
