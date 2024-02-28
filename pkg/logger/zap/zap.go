package zap

import (
	"log"

	"github.com/CesarDelgadoM/api-reports/config"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger(config *config.Config) {
	logzap, err := zap.NewDevelopment()
	if err != nil {
		log.Panic(err)
	}

	Logger = logzap.Sugar()
}
