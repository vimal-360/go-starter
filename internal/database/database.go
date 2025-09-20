package database

import (
	"fmt"
	"go-workflow-rnd/internal/config"
	"go-workflow-rnd/internal/models"
	"log"

	"github.com/samber/do/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode, cfg.DBTimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to PostgreSQL database: %s@%s:%s/%s", cfg.DBUser, cfg.DBHost, cfg.DBPort, cfg.DBName)
	return db, nil
}

func AutoMigrate(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}

func NewDatabaseService(injector do.Injector) (*gorm.DB, error) {
	cfg := do.MustInvoke[*config.Config](injector)
	db, err := Connect(cfg)
	if err != nil {
		return nil, err
	}

	if err := AutoMigrate(db, models.GetAllModels()...); err != nil {
		return nil, err
	}

	return db, nil
}