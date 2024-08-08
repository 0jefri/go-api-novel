package controller

import (
	"go-novel-api/model"
	"go-novel-api/usecase"
	"go-novel-api/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NovelController struct {
	router  *gin.Engine
	useCase usecase.NovelUsecase
}

func (e *NovelController) createHandler(c *gin.Context) {
	var novel model.Novel
	if err := c.ShouldBindJSON(&novel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	novel.Id = common.GenerateUUID()
	err := e.useCase.RegisterNewNovel(novel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Succes Create New Novel",
		"data":    novel,
	})
}

func NewNovelController(router *gin.Engine, nvlUseCase usecase.NovelUsecase) {
	ctr := &NovelController{
		router:  router,
		useCase: nvlUseCase,
	}

	routerGroup := ctr.router.Group("api/v1")
	routerGroup.POST("/novel", ctr.createHandler)
}
