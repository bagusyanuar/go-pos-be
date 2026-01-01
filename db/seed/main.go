package main

import (
	"github.com/bagusyanuar/go-pos-be/db/seed/seeder"
	"github.com/bagusyanuar/go-pos-be/internal/shared/config"
)

func main() {
	viper := config.NewViper()
	dbConfig := config.NewDatabaseConfig(viper)
	db := config.NewDatabaseConnection(dbConfig)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	seeder.Seed(db)

}
