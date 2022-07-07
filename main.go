package main

import (
	"github.com/chenhu1001/marketool/config"
	"github.com/chenhu1001/marketool/internal"
)

func main() {
	internal.Initialize()
	config.SetupRoutes()
	internal.Router.Run(":8080")
}
