package serializer

import (
	"gin_demo/models"
)

type Photo struct {
	Id        string `json:"id"`
	ImagePath string `json:"image_path"`
}

type Detail struct {
	Id     string  `json:"id"`
	Photos []Photo `json:"photos"`
	Title  string  `json:"title"`
	Price  string  `json:"price"`
}

// 序列化图片
func BuildPhotos() []Photo {
	var json_photo []Photo
	photos := models.GetPhotos()
	for _, v := range photos {
		json_photo = append(json_photo, Photo{
			Id:        v.Id,
			ImagePath: v.ImagePath,
		})
	}
	return json_photo
}

// 序列化标题
func BuildTitle() string {
	return models.GetTitle()
}

// 序列化价格
func BuildPrice() string {
	return models.GetPrice()
}

func BuildDetail(id string) Detail {
	return Detail{
		Id:     id,
		Photos: BuildPhotos(),
		Title:  BuildTitle(),
		Price:  BuildPrice(),
	}
}

func BuildDetailResponse(id string) Response {
	return Response{
		Msg:  "查询成功",
		Data: BuildDetail(id),
	}
}
