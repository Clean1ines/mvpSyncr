package session

import (
 "sync"
)

// State – этап взаимодействия пользователя с ботом.
type State int

const (
 StateStart State = iota
 StatePlatformSelection
 StateModeSelection
 StatePlaylistURL
 StateAuthorization
 StateSyncResult
)

// Session хранит состояние диалога с пользователем.
type Session struct {
 ChatID            int64
 State             State
 SelectedPlatforms map[string]bool // Например: {"Spotify": true, "YouTube Music": true}
 SelectedMode      string          // "AllLiked" или "Playlists"
 PlaylistURL       string
}

var sessions = make(map[int64]*Session)
var mutex sync.Mutex

// GetSession возвращает сессию пользователя или создаёт новую.
func GetSession(chatID int64) *Session {
 mutex.Lock()
 defer mutex.Unlock()
 if sess, ok := sessions[chatID]; ok {
  return sess
 }
 newSess := &Session{
  ChatID:            chatID,
  State:             StateStart,
  SelectedPlatforms: make(map[string]bool),
 }
 sessions[chatID] = newSess
 return newSess
}

// ResetSession сбрасывает сессию пользователя.
func ResetSession(chatID int64) {
 mutex.Lock()
 defer mutex.Unlock()
 sessions[chatID] = &Session{
  ChatID:            chatID,
  State:             StateStart,
  SelectedPlatforms: make(map[string]bool),
 }
}