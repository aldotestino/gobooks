package main

import (
	models "app/src/models"
)

func main() {
	db := models.NewPGBookDatabse()
	db.Migrate()
}
