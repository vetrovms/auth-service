package config

import "os"

// Env Об'єкт конфігурації зі змінних оточення та інші налаштування.
type Env struct {
	LogPath   string // Шлях до файлу логування.
	DbDsn     string // Доступ до бази даних.
	WebPort   string // Порт веб застосунку.
	SecretKey string // Таємний ключ для генерації jwt токена.
}

// logPath Назва змінної оточення що містить шлях до файлу логування.
const logPath = "AUTH_LOG_PATH"

// dbDsn Назва змінної оточення що містить доступ до бази даних.
const dbDsn = "AUTH_DB_DSN"

// webPort Назва змінної оточення що містить порт веб застосунку.
const webPort = "AUTH_WEB_PORT"

// secretKey Назва змінної оточення що містить таємний ключ для генерації jwt токена.
const secretKey = "AUTH_SECRET"

// NewEnv Повертає об'єкт конфігурації, заповнений зі змінних оточення.
func NewEnv() Env {
	return Env{
		LogPath:   os.Getenv(logPath),
		DbDsn:     os.Getenv(dbDsn),
		WebPort:   os.Getenv(webPort),
		SecretKey: os.Getenv(secretKey),
	}
}
