package repository

import (
	"context"
	"encoding/json"
	"go-redis/config"
	"go-redis/src/model"
)

type dataRepository struct {
	Cfg config.Config
}

func NewPostRepository(cfg config.Config) model.DataRepository {
	return &dataRepository{Cfg: cfg}
}

func (d *dataRepository) SetRedis(ctx context.Context, data *model.Data) (bool, error) {
	key := "setKey" + data.Key

	_, err := d.Cfg.Redis().Get(ctx, key)
	if err != nil {
		err := d.Cfg.Redis().Set(ctx, key, data.Value)
		if err != nil {
			return false ,err
		}

		return true, nil
	}

	return true, nil
}

func (d *dataRepository) GetRedis(ctx context.Context, key string) ([]*model.Data, error) {
	var dataFetch []*model.Data

	arrData, err := d.Cfg.Redis().Get(ctx, key)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(arrData), &dataFetch)

	if err != nil {
		return nil, err
	}

	return	dataFetch, nil
}
