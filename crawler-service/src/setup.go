package main

import (
	"crawler-service/src/component"
	"crawler-service/src/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createDsnDb(username, password, host, port, dbName string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)
}

func setupAppContext(config *config.AppConfig) component.AppContext {
	databaseDsn := createDsnDb(config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DatabaseName)

	FDDatabase, err := gorm.Open(postgres.Open(databaseDsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect database notification- ", err)
	}

	FDDatabase = FDDatabase.Debug()

	appCtx := component.NewAppContext(config, FDDatabase)

	return appCtx
}
