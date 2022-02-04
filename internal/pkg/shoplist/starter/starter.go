package starter

import (
	"github.com/gin-gonic/gin"
	"github.com/lgarcia93/shoplist/internal/db"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/controller"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/repository"
	"github.com/swaggo/files"       //
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func InitializeHandlers() {
	r := gin.Default()

	db, err := db.DbManagerImpl{}.NewConnection()

	defer db.Close()

	if err != nil {
		log.Fatalln("error opening connection")
	}

	c := controller.ShopListControllerImpl{
		Repository: repository.NewShopListRepository(db),
	}

	r.Group("/api/v1")
	{
		r.POST("/shopitem", c.Create)
		r.PUT("/shopitem", c.Update)
		r.DELETE("/shopitem/:id", c.Delete)
		r.GET("/shopitem", c.GetAll)
		r.GET("/shopitem/:id", c.Get)

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run() // listen and serve o
}
