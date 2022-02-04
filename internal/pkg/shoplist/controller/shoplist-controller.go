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

// @Summary      Create a ShopItem
// @Description  Returns the new  {object} model.ShopItem that was created
// @Tags         shopitem
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ShopItemID"
// @Success      200  {object}  model.ShopItem
// @Failure		 500
// @Failure 	 404
// @Router       /shopitem/{id} [post]
func (s ShopListControllerImpl) Create(ctx *gin.Context) {
	var shopItem model.ShopItem

	ctx.BindJSON(&shopItem)

	affectedRows, err := s.Repository.Create(shopItem)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if affectedRows > 0 {
		ctx.JSON(http.StatusAccepted, shopItem)
	}
}

// @Summary      Updates a ShopItem
// @Description  Updates a  ShopItem with the supplied ShopItem object
// @Tags         shopitem
// @Accept       json
// @Produce      json
// @Param        shopItem  body model.ShopItem  true  "ShopItemID"
// @Success      200  {object}  model.ShopItem
// @Failure		 500
// @Failure 	 404
// @Router       /shopitem/ [put]
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

// @Summary      Deletes a ShopItem
// @Description  Deletes a  ShopItem with the supplied id value
// @Tags         shopitem
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ShopItemID"
// @Success      200  {object}  model.ShopItem
// @Failure		 500
// @Failure 	 404
// @Router       /shopitem/{id} [delete]
func (s ShopListControllerImpl) Delete(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	affectedRows, err := s.Repository.Delete(id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	if affectedRows > 0 {
		ctx.Status(http.StatusAccepted)

		return
	}

	ctx.Status(http.StatusNotModified)
}

// @Summary      Retrieves a ShopItem
// @Description  Retrieves a ShopItem with the supplied id value
// @Tags         shopitem
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ShopItemID"
// @Success      200  {object}  model.ShopItem
// @Failure		 500
// @Failure 	 404
// @Router       /shopitem/{id} [get]
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

	if shopItem != nil {
		ctx.JSON(http.StatusAccepted, shopItem)
		return
	}

	ctx.Status(http.StatusNotFound)

}

// @Summary      Retrieves all the shopitems
// @Description  Retrieves all the shopitems
// @Tags         shopitem
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.ShopItem
// @Failure		 500
// @Router       /shopitem/ [get]
func (s ShopListControllerImpl) GetAll(ctx *gin.Context) {

	items, err := s.Repository.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error"))

		return
	}

	ctx.JSON(200, items)
}
