package main

import (
	"fmt"
	"github.com/chenhu1001/marketool/config"
	"github.com/chenhu1001/marketool/goutils"
	"github.com/chenhu1001/marketool/internal"
)

func main() {
	ip, err := goutils.GetLocalIP()
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(ip)
	}

	internal.Initialize()
	config.SetupRoutes()
	internal.Router.Run(":8080")
}
