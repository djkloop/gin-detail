package service

import (
	"gin_demo/serializer"
)

// GetDetail 获取detail数据
func GetDetail(id string) serializer.Response {
	return serializer.BuildDetailResponse(id)
}
