package telegramclient

import (
 "log"
 "telegram-sync-bot/config"

 tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Client инкапсулирует объект бота.
type Client struct {
 Bot *tgbotapi.BotAPI
}

// NewClient инициализирует Telegram-бота с использованием конфигурации.
func NewClient(cfg *config.Config) *Client {
 bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
 if err != nil {
  log.Panic(err)
 }
 // Выключаем режим отладки для продакшена
 bot.Debug = false

 return &Client{Bot: bot}
}