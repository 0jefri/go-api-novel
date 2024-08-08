package delivery

import (
	"fmt"
	"go-novel-api/config"
	"go-novel-api/delivery/controller"
	"go-novel-api/repository"
	"go-novel-api/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	nvlUseCase usecase.NovelUsecase
	engine     *gin.Engine
	host       string
}

func (a *appServer) initController() {
	controller.NewNovelController(a.engine, a.nvlUseCase)
}

func (a *appServer) Run() {
	a.initController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err.Error())
	}
}

func Server() *appServer {
	engine := gin.Default()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}
	dbConn, _ := config.NewDbConnection(cfg)
	nvlRepo := repository.NewNovelRepository(dbConn.Conn())
	nvlUseCase := usecase.NewNovelUsecase(nvlRepo)
	host := fmt.Sprintf("%s : %s", cfg.ApiHost, cfg.ApiPort)

	return &appServer{
		engine:     engine,
		nvlUseCase: nvlUseCase,
		host:       host,
	}
}
