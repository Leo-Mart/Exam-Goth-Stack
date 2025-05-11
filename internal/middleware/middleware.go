package middleware

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type Middleware struct {
	session *scs.SessionManager
}

func NewMiddleware(session *scs.SessionManager) (*Middleware, error) {
	return &Middleware{
		session: session,
	}, nil
}

func (m *Middleware) SessionLoad(next http.Handler) http.Handler {
	return m.session.LoadAndSave(next)
}
