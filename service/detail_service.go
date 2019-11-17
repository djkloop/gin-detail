package service

import (
	"gin_demo/serializer"
)

func GetDetail(id string) serializer.Response {
	return serializer.BuildDetailResponse(id)
}
