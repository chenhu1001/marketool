package goutils

import (
	"log"
	"net/http"
	"time"
)

// Bark推送
func Push(str string) {
	now := time.Now()
	currentDate := now.Format("2006-01-02")
	_, err := http.Get("https://api.day.app/Y3uKSZF6URZQTU7FXuTUUM/" + "【" + currentDate + "】" + str)
	if err != nil {
		log.Println("err")
	}
}
