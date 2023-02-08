package src

import (
	"fmt"
	"go-redis/config"
	"go-redis/src/delivery"
	"go-redis/src/helper/validator"
	"go-redis/src/repository"
	"go-redis/src/service"
	"log"
	"net/http"

	validatorEngine "github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	server struct {
		httpServer *echo.Echo
		cfg        config.Config
	}

	Server interface {
		Run()
	}
)

func InitServer(cfg config.Config) Server {
	e := echo.New()
	e.HideBanner = true
	e.Validator = &validator.GoPlaygroundValidator{
		Validator: validatorEngine.New(),
	}

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &server{
		httpServer: e,
		cfg:        cfg,
	}

}

func (s *server) Run() {
	s.httpServer.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "success",
			"message": "Hello, world!" + s.cfg.ServiceName() + " " + s.cfg.ServiceEnvirontment(),
		})
	})

	dataRepository := repository.NewPostRepository(s.cfg)
	dataService := service.NewDataService(dataRepository)
	dataDelivery := delivery.NewDataDelivery(dataService)
	dataGroup := s.httpServer.Group("/data")
	dataDelivery.Mount(dataGroup)

	err := s.httpServer.Start(fmt.Sprintf(":%d", s.cfg.ServicePort()));
	if err != nil {
		log.Panic(err)
	}
}
