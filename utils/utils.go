package utils

import (
	"go.uber.org/zap"
)

//var Zapper *zap.Logger

func InitZap(cfg zap.Config) *zap.Logger {
	//var cfg zap.Config

	//Zapper, _ = zap.NewProduction()
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	//defer Zapper.Sync()
	defer logger.Sync()

	logger.Info("logger construiction succeeed...")
	return logger
}
