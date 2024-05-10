package utility

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

// Connect to Database
func DbConnect() {
	// Data source name config
	dsn := "host=localhost user=postgres password=Yogesh@123 dbname=GoDB search_path=blogging_db port=5432 sslmode=disable"
	// Db connection
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	sql, err := Db.DB()

	if err != nil {
		panic(err.Error())
	}

	if err := sql.Ping(); err != nil {
		panic(err.Error())
	}
}
