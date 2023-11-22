package migrate

import (
	"github.com/danielalejandrorosero/jwt_gin/initialize"
	"github.com/danielalejandrorosero/jwt_gin/models"
)

func init() {
	initialize.LoadEnv()
	initialize.DataBase()
}

func main() {
	initialize.DB.AutoMigrate(&models.User{})
}
