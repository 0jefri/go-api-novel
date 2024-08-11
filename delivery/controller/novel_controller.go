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

func (e *NovelController) listHandler(c *gin.Context) {
	// judul := c.Query("judul")
	// penerbit := c.Query("penerbit")
	// tahunTerbit := c.Query("tahunTerbit")
	// penulis := c.Query("penulis")
	novels, err := e.useCase.FindAllNovels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	status := map[string]any{
		"code":    200,
		"message": "Get All Data Successfully",
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   novels,
	})
}

func (e *NovelController) getHandler(c *gin.Context) {
	id := c.Param("id")
	novel, err := e.useCase.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Get Novel By Id",
		"data":    novel,
	})
	return
}

func (e *NovelController) UpdateHandler(c *gin.Context) {
	var novel model.Novel
	if err := c.ShouldBindJSON(&novel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Update Novel",
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
	routerGroup.GET("/novel", ctr.listHandler)
	routerGroup.GET("/novel/:id", ctr.getHandler)
	routerGroup.PUT("/novel", ctr.UpdateHandler)
}
