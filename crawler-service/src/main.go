package main

import (

	// "github.com/gin-contrib/cors"

	"crawler-service/src/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	appConfig, err := config.NewAppConfig("./config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	if appConfig.Server.ModeRun == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	appCtx := setupAppContext(appConfig)

	r := gin.Default()
	// r.Use(cors.Default())

	setupRouter(r, appCtx)

	// r.Run(":", appConfig.Server.Port)
	fmt.Println("Listening on: " + "http://" + appConfig.Server.Host + ":" + appConfig.Server.Port)
	r.Run(":" + appConfig.Server.Port)

}
