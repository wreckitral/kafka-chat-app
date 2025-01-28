package main

import (
	"github.com/wreckitral/kafka-chat-app/cmd/server"
	"github.com/wreckitral/kafka-chat-app/pkg/config"
)

// @title kafka Chat app
// @version 1.0
// @description Chat app integrated with kafka, made using Golang and fiber
// @contact.name Defhanaya Sofhiea
// @contact.email defhanayasofhiea@gmail.com
// @termsOfService
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:4000
// @BasePath /api
func main() {
    config.LoadAllConfigs(".env")

    server.Serve()
}
