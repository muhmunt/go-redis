package service

import (
	"context"
	"go-redis/src/model"
	"go-redis/src/request"
	"go-redis/src/response"
)

type dataService struct {
	dataRepository model.DataRepository
}

func NewDataService(data model.DataRepository) model.DataService {
	return &dataService{}
}

func (d *dataService) GetData(ctx context.Context, key string) (*response.DataResponse, error) {
	arrData ,err := d.dataRepository.GetRedis(ctx, key)

	if err != nil {
		return nil, err
	}

	response := new(response.DataResponse)

	response.Value = arrData
	response.Meta.Key = key

	return response, nil
}

func (d *dataService) SetData(ctx context.Context, request request.DataRequest) (*response.DataResponse, error) {
	newData := model.Data{
		Key: request.Key,
		Value: request.Value,
	}

	_, err := d.dataRepository.SetRedis(ctx, &newData)

	if err != nil {
		return nil, err
	}

	response := new(response.DataResponse)
	response.Meta.Key = newData.Key

	return response, nil
}
