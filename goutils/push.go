package goutils

import (
	"log"
	"net/http"
)

// Bark推送
func Push(str string) {
	_, err := http.Get("https://api.day.app/Y3uKSZF6URZQTU7FXuTUUM/" + str)
	if err != nil {
		log.Println("err")
	}
}
