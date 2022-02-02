package starter

import (
	"github.com/gin-gonic/gin"
	"github.com/lgarcia93/shoplist/internal/db"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/controller"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/repository"
	"log"
)

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

	r.POST("/shoplist", c.Create)
	r.PUT("/shoplist", c.Update)
	r.DELETE("/shoplist/:id", c.Delete)
	r.GET("/shoplist", c.GetAll)
	r.GET("/shoplist/:id", c.Get)

	r.Run() // listen and serve o
}
