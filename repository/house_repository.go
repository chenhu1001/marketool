package repository

import (
	"github.com/chenhu1001/marketool/models"
)

type HouseRepository struct {
	//Repository
	//model models.House
}

func (r *HouseRepository) FetchAll() []models.House {
	houses := []models.House{}
	//Database.Find(&houses)
	var house models.House
	house.ID = 3
	house.Title = "Ejidal"
	house.Content = "felis ut at"
	house.Address = "lacinia aenean sit"
	houses = append(houses, house)

	return houses
}

func (r *HouseRepository) FetchById(id int) *models.House {
	var house models.House
	//Database.First(&house, id)
	house.ID = 2
	house.Title = "San Andrés"
	house.Content = "vel nulla eget"
	house.Address = "quam suspendisse potenti nullam"

	return &house
}

func (r *HouseRepository) FindByTitle(title string) []models.House {
	// sqlStatement := fmt.Sprintf(`SELECT * FROM public.actor WHERE first_name ILIKE '%%%s%%'`, title)

	//rows := QueryDB(sqlStatement)
	//houseList := buildHouse(rows)
	houses := []models.House{}
	//Database.Where("title ILIKE ?", "%"+title+"%").Find(&houses)

	var house models.House
	house.ID = 1
	house.Title = "Xiachengzi"
	house.Content = "turpis elementum ligula vehicula"
	house.Address = "bibendum morbi non quam"
	houses = append(houses, house)

	return houses
}
