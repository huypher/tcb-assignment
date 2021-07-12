// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package app

import (
	"context"
	"tcb-assignment/internal/infra"
)

// Injectors from wire.go:

func InitApplication(ctx context.Context) (*ApplicationContext, func(), error) {
	appConfig, err := infra.ProvideConfig()
	if err != nil {
		return nil, nil, err
	}
	service := infra.ProvideAuthService(appConfig)
	cache := infra.ProvideCacheService()
	poolRepo := infra.ProvidePoolRepo()
	producer := infra.ProvidePoolProducer(appConfig)
	poolsService := infra.ProvidePoolService(cache, poolRepo, producer)
	restAPIHandler := infra.ProvideRestAPIHandler(service, poolsService)
	restService, cleanup, err := infra.ProvideRestService(appConfig, restAPIHandler)
	if err != nil {
		return nil, nil, err
	}
	poolConsumer := infra.ProvidePoolConsumer(poolsService, appConfig)
	applicationContext := &ApplicationContext{
		ctx:          ctx,
		cfg:          appConfig,
		restSrv:      restService,
		poolProducer: producer,
		poolConsumer: poolConsumer,
	}
	return applicationContext, func() {
		cleanup()
	}, nil
}
