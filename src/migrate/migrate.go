package migrate

import (
	"log"

	"ufc.com/deti/go-dad/src/database"
	"ufc.com/deti/go-dad/src/model"
)

// Migrate -- Migrate
func Migrate() {

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.Book{})

	db.Close()
}
