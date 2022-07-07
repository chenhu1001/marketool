// handlers.houses.go

package handlers

import (
	"github.com/chenhu1001/marketool/internal"
	"github.com/chenhu1001/marketool/models"
	"github.com/chenhu1001/marketool/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	houseRepository := repository.HouseRepository{}

	title := c.Query("title")
	houses := []models.House{}

	if len(title) > 0 {
		houses = houseRepository.FindByTitle(title)
	} else {
		houses = houseRepository.FetchAll()
	}

	internal.Render(c, gin.H{
		"title":   "Home Page",
		"payload": houses},
		"index.html")
}

func GetHomeHouse(c *gin.Context) {
	houseRepository := repository.HouseRepository{}

	house := houseRepository.FetchById(1)

	if house == nil {
		return
	}

	internal.Render(c, gin.H{
		"title":   house.Title,
		"payload": house}, "house.html")
}

func GetHouse(c *gin.Context) {
	houseRepository := repository.HouseRepository{}

	houseId, err := strconv.Atoi(c.Param("house_id"))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	house := houseRepository.FetchById(houseId)

	if house == nil {
		// If the article is not found, abort with an error
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	internal.Render(c, gin.H{
		"title":   house.Title,
		"payload": house}, "house.html")
}
