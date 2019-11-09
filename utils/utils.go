package utils

import "go.uber.org/zap"

var Zapper *zap.Logger

func InitZap() {
	var cfg zap.Config

	Zapper, _ = zap.NewProduction()
	defer Zapper.Sync()
}
