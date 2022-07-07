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

	return houses
}

func (r *HouseRepository) FetchById(id int) *models.House {
	var house models.House
	//Database.First(&house, id)
	return &house
}

func (r *HouseRepository) FindByTitle(title string) []models.House {
	// sqlStatement := fmt.Sprintf(`SELECT * FROM public.actor WHERE first_name ILIKE '%%%s%%'`, title)

	//rows := QueryDB(sqlStatement)
	//houseList := buildHouse(rows)
	houses := []models.House{}
	//Database.Where("title ILIKE ?", "%"+title+"%").Find(&houses)

	return houses
}
