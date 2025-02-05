package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"teca_notifications/api" // Importa el paquete api
)

func Connect() (*gorm.DB, error) {
	dsn := "host=dpg-cufsuran91rc73cjgi7g-a.oregon-postgres.render.com user=teca_notifications_user password=huPnZFhBQgH8HwV5z4qfwVsyIDFxBllP dbname=teca_notifications port=5432 sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrar los modelos
	err = db.AutoMigrate(&api.User{}, &api.Task{}, &api.Comment{}, &api.Activity{}, &api.Notification{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
