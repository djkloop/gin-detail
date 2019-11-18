package serializer

import (
	"gin_demo/models"
)

// Photo 注释
type Photo struct {
	ID        string `json:"id"`
	ImagePath string `json:"image_path"`
}

// Detail 注释
type Detail struct {
	ID     string  `json:"id"`
	Photos []Photo `json:"photos"`
	Title  string  `json:"title"`
	Price  string  `json:"price"`
}

// BuildPhotos 序列化图片
func BuildPhotos() []Photo {
	var jsonPhoto []Photo
	photos := models.GetPhotos()
	for _, v := range photos {
		jsonPhoto = append(jsonPhoto, Photo{
			ID:        v.Id,
			ImagePath: v.ImagePath,
		})
	}
	return jsonPhoto
}

// BuildTitle 序列化标题
func BuildTitle() string {
	return models.GetTitle()
}

// BuildPrice 序列化价格
func BuildPrice() string {
	return models.GetPrice()
}

// BuildDetail 序列化返回结果
func BuildDetail(id string) Detail {
	return Detail{
		ID:     id,
		Photos: BuildPhotos(),
		Title:  BuildTitle(),
		Price:  BuildPrice(),
	}
}

// BuildDetailResponse 返回结果
func BuildDetailResponse(id string) Response {
	return Response{
		Msg:  "查询成功",
		Data: BuildDetail(id),
	}
}
