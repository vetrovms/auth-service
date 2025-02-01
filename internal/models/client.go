package models

import "gorm.io/gorm"

// Client Модель клієнта (таблиця де зберігаються client_id та client_secret сервісів-клієнтів).
type Client struct {
	gorm.Model
	ClientId     string `form:"client_id" validate:"required,max=255"`
	ClientSecret string `form:"client_secret" validate:"required,max=255"`
}
