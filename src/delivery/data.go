package delivery

import (
	"go-redis/src/model"
	"go-redis/src/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type dataDelivery struct {
	dataService model.DataService
}

type DataDelivery interface {
	Mount(group *echo.Group)
}

func NewDataDelivery(dataService model.DataService) DataDelivery {
	return &dataDelivery{dataService: dataService}
}

func (d *dataDelivery) Mount(group *echo.Group) {
	group.GET("/fetch", d.FetchHandler)
	group.POST("/store", d.FetchHandler)
}

func (d *dataDelivery) FetchHandler(c echo.Context) error {
	ctx := c.Request().Context()

	key := c.QueryParam("key")

	data, err := d.dataService.GetData(ctx, key)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (d *dataDelivery) StoreHandle(c echo.Context) error {
	var req request.DataRequest
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		return err
	}

	err := c.Validate(req)
	if err != nil {
		return err
	}

	_, err = d.dataService.SetData(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
