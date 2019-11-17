package models

import "fmt"

type photo struct {
	Id        string
	ImagePath string
}

func GetPhotos() []photo {

	photos := []photo{
		{
			Id:        "pic1",
			ImagePath: "img01.jpg",
		},
		{
			Id:        "pic2",
			ImagePath: "img02.jpg",
		},
		{
			Id:        "pic3",
			ImagePath: "img03.jpg",
		},
		{
			Id:        "pic4",
			ImagePath: "img04.jpg",
		},
		{
			Id:        "pic5",
			ImagePath: "img05.jpg",
		},
		{
			Id:        "pic6",
			ImagePath: "img06.jpg",
		},
	}
	return photos
}

func GetTitle() string {
	return `golang gin测试标题`
}

func GetPrice() string {
	price := 139900
	intPrice := price / 100
	decPrice := price % 100
	return fmt.Sprintf("%d.%02d", intPrice, decPrice)
}
