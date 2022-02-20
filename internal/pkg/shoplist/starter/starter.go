package starter

import (
	"github.com/gin-gonic/gin"
	"github.com/lgarcia93/shoplist/internal/db"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/controller"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/repository"
	"github.com/swaggo/files"       //
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
	"os"
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

	env := getAppEnv()

	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := db.DbManagerImpl{}.NewConnection(env)

	defer db.Close()

	if err != nil {
		log.Fatalln("error opening connection")
	}

	c := controller.ShopListControllerImpl{
		Repository: repository.NewShopListRepository(db),
	}

	g := r.Group("/api/v1")
	{
		g.POST("/shopitem", c.Create)
		g.PUT("/shopitem", c.Update)
		g.DELETE("/shopitem/:id", c.Delete)
		g.GET("/shopitem", c.GetAll)
		g.GET("/shopitem/:id", c.Get)

		if getAppEnv() == "dev" {
			g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}

	r.Run(":5000") // listen and serve o
}

func getAppEnv() string {
	env := os.Getenv("shoplist_api_env")

	if env == "" {
		env = "dev"
	}

	return env
}
