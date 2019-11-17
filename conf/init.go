package conf

import (
	"gin_demo/models"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() {
	godotenv.Load()
	models.InitDB(os.Getenv("MYSQL_DSN"))
}
