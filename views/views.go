package views

import (
 tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// StartView – стартовый экран с кнопкой «Старт».
func StartView(chatID int64, text string) tgbotapi.MessageConfig {
 msg := tgbotapi.NewMessage(chatID, text)
 msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
  tgbotapi.NewInlineKeyboardRow(
   tgbotapi.NewInlineKeyboardButtonData("Старт", "start"),
  ),
 )
 return msg
}

// PlatformSelectionView – экран выбора платформ с кнопками и кнопкой «Далее».
func PlatformSelectionView(chatID int64, text string) tgbotapi.MessageConfig {
 btnSpotify := tgbotapi.NewInlineKeyboardButtonData("Spotify", "platform_spotify")
 btnYoutube := tgbotapi.NewInlineKeyboardButtonData("YouTube Music", "platform_youtube")
 nextBtn := tgbotapi.NewInlineKeyboardButtonData("Далее", "next")
 keyboard := tgbotapi.NewInlineKeyboardMarkup(
  tgbotapi.NewInlineKeyboardRow(btnSpotify, btnYoutube),
  tgbotapi.NewInlineKeyboardRow(nextBtn),
 )
 msg := tgbotapi.NewMessage(chatID, text)
 msg.ReplyMarkup = keyboard
 return msg
}

// ModeSelectionView – экран выбора режима синхронизации.
func ModeSelectionView(chatID int64, text string) tgbotapi.MessageConfig {
 btnAllLiked := tgbotapi.NewInlineKeyboardButtonData("AllLiked", "mode_allliked")
 btnPlaylists := tgbotapi.NewInlineKeyboardButtonData("Playlists", "mode_playlists")
 keyboard := tgbotapi.NewInlineKeyboardMarkup(
  tgbotapi.NewInlineKeyboardRow(btnAllLiked, btnPlaylists),
 )
 msg := tgbotapi.NewMessage(chatID, text)
 msg.ReplyMarkup = keyboard
 return msg
}

// AuthorizationView – имитация экрана авторизации (в реальном проекте здесь откроется веб-вью).
func AuthorizationView(chatID int64, text string) tgbotapi.MessageConfig {
 authText := text + "\n[Здесь будет авторизация в браузере]"
 msg := tgbotapi.NewMessage(chatID, authText)
 msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
  tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Рестарт", "restart")),
 )
 return msg
}

// TextMessage – простое текстовое сообщение.
func TextMessage(chatID int64, text string) tgbotapi.MessageConfig {
 return tgbotapi.NewMessage(chatID, text)
}