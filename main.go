package main

import (
	"fmt"
	"github.com/chenhu1001/marketool/goutils"
)

func main() {
	ip, err := goutils.GetLocalIP()
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(ip)
	}
}
