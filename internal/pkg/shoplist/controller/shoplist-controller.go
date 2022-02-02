package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/model"
	"github.com/lgarcia93/shoplist/internal/pkg/shoplist/repository"
	"net/http"
	"strconv"
)

type ShopListController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type ShopListControllerImpl struct {
	Repository repository.ShopListRepository
}

func (s ShopListControllerImpl) Create(ctx *gin.Context) {

	var shopItem model.ShopItem

	ctx.BindJSON(&shopItem)

	affectedRows, err := s.Repository.Create(shopItem)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	if affectedRows > 0 {
		ctx.JSON(http.StatusAccepted, shopItem)
	}
}

func (s ShopListControllerImpl) Update(ctx *gin.Context) {
	var shopItem model.ShopItem

	ctx.BindJSON(&shopItem)

	affectedRows, err := s.Repository.Update(shopItem)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	if affectedRows > 0 {
		ctx.JSON(http.StatusAccepted, shopItem)
	}
}

func (s ShopListControllerImpl) Delete(ctx *gin.Context) {
	var shopItem model.ShopItem

	ctx.BindJSON(&shopItem)

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	affectedRows, err := s.Repository.Delete(id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	if affectedRows > 0 {
		ctx.JSON(http.StatusAccepted, shopItem)
	}
}

func (s ShopListControllerImpl) Get(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}


	shopItem, err := s.Repository.Get(id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusAccepted, shopItem)
}

func (s ShopListControllerImpl) GetAll(ctx *gin.Context) {

	items, err := s.Repository.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error"))

		return
	}

	ctx.JSON(200, items)

}

