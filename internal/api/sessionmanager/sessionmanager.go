package sessionmanager

import (
	"time"

	"github.com/alexedwards/scs/v2"
)

func New() *scs.SessionManager {
	// TODO: Add Redis store for sessions
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Name = "session_id"

	return sessionManager
}
