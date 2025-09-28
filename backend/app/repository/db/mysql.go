package dbrepository

import (
	"be-tasking/app/repository"
	"be-tasking/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pressly/goose"
)

type mySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) repository.MySQLRepoInterface {
	return &mySQLRepository{db}
}

// create new MySQL database connection with gorm
func NewMySQLConn(conf config.MySQLConfiguration) (*gorm.DB, error) {
	connURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s&loc=%s",
		conf.DBUser,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
		conf.DBOptions,
		conf.Locale)

	db, err := gorm.Open("mysql", connURL)
	if err != nil {
		return nil, err
	}
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.LogMode(true)

	log.Println("database connected!")

	// Auto migrate
	err = goose.Up(db.DB(), "app/repository/db/migration")
	if err != nil {
		fmt.Println("Error when running migration:", err.Error())
	}

	return db, nil
}
