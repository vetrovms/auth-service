package logger

import (
	"os"
	"sync"
	"testing"

	"auth/internal/config"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Log Повертає компонент логування.
func Log() *logrus.Logger {
	var once sync.Once
	once.Do(func() {
		env := config.NewEnv()
		log.SetFormatter(&logrus.JSONFormatter{})
		if !testing.Testing() {
			file, err := os.OpenFile(env.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Fatal(err)
			}
			log.SetOutput(file)
		}
	})
	return log
}
