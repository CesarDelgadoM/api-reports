package database

import (
	"fmt"

	"github.com/CesarDelgadoM/api-reports/config"
	"github.com/CesarDelgadoM/api-reports/pkg/logger/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	*gorm.DB
}

func ConnectPostgresDB(config config.PostgresConfig) *PostgresDB {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		zap.Logger.Fatal("Connect to postgres db failed: ", err)
	}
	zap.Logger.Info("Connect to postgres db success")

	return &PostgresDB{db}
}
