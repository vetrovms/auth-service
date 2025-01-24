package connection

import (
	"auth/internal/config"
	"auth/internal/logger"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Db Повертає з'єднання з базою даних.
func Db() *gorm.DB {
	var once sync.Once
	once.Do(func() {
		env := config.NewEnv()
		var err error
		db, err = gorm.Open(postgres.Open(env.DbDsn))
		if err != nil {
			logger.Log().Fatal("failed to connect database")
		}
	})
	return db
}
