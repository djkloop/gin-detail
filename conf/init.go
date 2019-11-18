package conf

import (
	"gin_demo/models"
	"os"

	"github.com/joho/godotenv"
)

// InitConfig 初始化配置
func InitConfig() {
	godotenv.Load()
	models.InitDB(os.Getenv("MYSQL_DSN"))
}
