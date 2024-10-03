package sessionmanager

import (
	"encoding/gob"
	"time"

	"github.com/ackuq/wishlist-backend/internal/api/auth"
	"github.com/alexedwards/scs/v2"
)

var SessionManager *scs.SessionManager

func Init() {
	// TODO: Add Redis store for sessions

	// Register custom structs
	gob.Register(auth.Claims{})

	// Create session manager
	SessionManager = scs.New()
	SessionManager.Lifetime = 24 * time.Hour
	SessionManager.Cookie.Name = "session_id"
}
