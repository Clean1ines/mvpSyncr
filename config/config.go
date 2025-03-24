package config

import (
 "os"
)

// Config хранит параметры конфигурации приложения.
type Config struct {
 TelegramToken string
 // При необходимости добавьте другие токены (Spotify, YouTube и т.п.)
}

// LoadConfig загружает конфигурацию из переменных окружения.
func LoadConfig() *Config {
 return &Config{
  TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
 }
}