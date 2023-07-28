package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "github.com/sreenathsvrm/voix-me/pkg/config"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp" + "(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?" + "parseTime=true&loc=Local"
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	return db, dbErr

}
