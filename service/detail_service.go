package service

import "gin_demo/models"

func GetPhotosService() []models.Photo {
	photos := models.GetPhotos()
	return photos
}

func GetPriceService() string {
	price := models.GetPrice()
	return price
}

func GetTitleService() string {
	title := models.GetTitle()
	return title
}
