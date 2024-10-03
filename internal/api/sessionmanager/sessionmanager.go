package sessionmanager

import (
	"encoding/gob"
	"time"

	"github.com/ackuq/wishlist-backend/internal/api/auth"
	"github.com/alexedwards/scs/v2"
)

func New() *scs.SessionManager {
	// TODO: Add Redis store for sessions

	// Register custom structs
	gob.Register(auth.Claims{})

	// Create session manager
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Name = "session_id"

	return sessionManager
}
