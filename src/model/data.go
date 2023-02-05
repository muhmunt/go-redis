package model

import (
	"context"
	"go-redis/src/request"
	"go-redis/src/response"
)

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DataRepository interface {
	SetRedis(ctx context.Context, data *Data) (bool, error)
	GetRedis(ctx context.Context, key string) ([]*Data, error)
}

type DataService interface {
	GetData(ctx context.Context, key string) (*response.DataResponse, error)
	SetData(ctx context.Context, request request.DataRequest) (*response.DataResponse, error)
}
