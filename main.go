package main

import (
 "log"
 "telegram-sync-bot/config"
 "telegram-sync-bot/telegramclient"
 "telegram-sync-bot/session"
 "telegram-sync-bot/handlers"

 tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
 // Загружаем конфигурацию из переменных окружения.
 cfg := config.LoadConfig()
 if cfg.TelegramToken == "" {
  log.Fatal("TELEGRAM_TOKEN не задан в окружении")
 }

 // Инициализируем Telegram-бота.
 client := telegramclient.NewClient(cfg)

 // Настраиваем получение обновлений.
 u := tgbotapi.NewUpdate(0)
 u.Timeout = 60
 updates := client.Bot.GetUpdatesChan(u)

 log.Println("Бот запущен...")

 // Отправляем стартовое сообщение первому пользователю (если нужно, можно инициировать по запросу)
 // Пример: client.Bot.Send(views.StartView(chatID, "Добро пожаловать! Нажмите 'Старт'."))

 // Основной цикл обработки обновлений.
 for update := range updates {
  var chatID int64
  if update.Message != nil {
   chatID = update.Message.Chat.ID
  } else if update.CallbackQuery != nil {
   chatID = update.CallbackQuery.Message.Chat.ID
  }
  if chatID == 0 {
   continue
  }
  sess := session.GetSession(chatID)

  msg, err := handlers.HandleUpdate(update, sess)
  if err != nil {
   log.Println("Ошибка обработки обновления:", err)
   continue
  }
  if msg != nil {
   _, err := client.Bot.Send(msg)
   if err != nil {
    log.Println("Ошибка отправки сообщения:", err)
   }
  }
 }
}