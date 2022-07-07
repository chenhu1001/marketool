package main

import (
	"github.com/chenhu1001/marketool/config"
	"github.com/chenhu1001/marketool/internal"
	"github.com/chenhu1001/marketool/repository"
)

func main() {
	internal.Initialize()
	config.SetupRoutes()
	repository.SetupDatabase()
	internal.Router.Run(":8080")
}
