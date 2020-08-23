package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database -- Struct
type Database struct {
	dialect  string
	host     string
	port     string
	user     string
	dbname   string
	password string
}

// New -- Cria uma nova Struct com os valores de configuração
func (d *Database) New() {
	d.dialect = "postgres"
	d.host = "godaddb.chct2onrz0xz.us-east-1.rds.amazonaws.com"
	d.port = "5432"
	d.user = "godad"
	d.password = "godad1234"
	d.dbname = "postgres"

}

// NewConnection -- Cria uma nova conexão com o banco
func NewConnection() (*gorm.DB, error) {
	databaseParams := Database{}
	databaseParams.New()
	stringConnection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		databaseParams.host, databaseParams.port,
		databaseParams.user, databaseParams.dbname,
		databaseParams.password)
	db, err := gorm.Open("postgres", stringConnection)
	return db, err
}
