package handlers

import (
 "telegram-sync-bot/session"
 "telegram-sync-bot/views"

 tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleUpdate обрабатывает обновления (callback и текстовые сообщения).
func HandleUpdate(update tgbotapi.Update, sess *session.Session) (tgbotapi.Chattable, error) {
 if update.CallbackQuery != nil {
  return handleCallbackQuery(update.CallbackQuery, sess)
 }
 if update.Message != nil {
  return handleMessage(update.Message, sess)
 }
 return nil, nil
}

func handleCallbackQuery(query *tgbotapi.CallbackQuery, sess *session.Session) (tgbotapi.Chattable, error) {
 data := query.Data
 switch data {
 case "start":
  sess.State = session.StatePlatformSelection
  return views.PlatformSelectionView(query.Message.Chat.ID, "Выберите платформы для синхронизации:"), nil
 case "platform_spotify":
  sess.SelectedPlatforms["Spotify"] = !sess.SelectedPlatforms["Spotify"]
  return views.PlatformSelectionView(query.Message.Chat.ID, "Выберите платформы для синхронизации:"), nil
 case "platform_youtube":
  sess.SelectedPlatforms["YouTube Music"] = !sess.SelectedPlatforms["YouTube Music"]
  return views.PlatformSelectionView(query.Message.Chat.ID, "Выберите платформы для синхронизации:"), nil
 case "next":
  sess.State = session.StateModeSelection
  return views.ModeSelectionView(query.Message.Chat.ID, "Выберите режим синхронизации:"), nil
 case "mode_allliked":
  sess.SelectedMode = "AllLiked"
  sess.State = session.StateAuthorization
  return views.AuthorizationView(query.Message.Chat.ID, "Авторизация для AllLiked: Открывается веб-вью..."), nil
 case "mode_playlists":
  sess.SelectedMode = "Playlists"
  sess.State = session.StatePlaylistURL
  return views.TextMessage(query.Message.Chat.ID, "Введите URL исходного плейлиста:"), nil
 case "restart":
  session.ResetSession(query.Message.Chat.ID)
  return views.StartView(query.Message.Chat.ID, "Стартовый экран. Нажмите 'Старт' чтобы начать снова."), nil
 default:
  return views.TextMessage(query.Message.Chat.ID, "Неизвестная команда: "+data), nil
 }
}

func handleMessage(message *tgbotapi.Message, sess *session.Session) (tgbotapi.Chattable, error) {
 // Если ожидается ввод URL плейлиста.
 if sess.State == session.StatePlaylistURL {
  sess.PlaylistURL = message.Text
  sess.State = session.StateAuthorization
  return views.AuthorizationView(message.Chat.ID, "URL получен. Открывается веб-вью для авторизации..."), nil
 }
 // По умолчанию – напоминание нажать кнопку.
 return views.TextMessage(message.Chat.ID, "Нажмите кнопку для продолжения."), nil
}