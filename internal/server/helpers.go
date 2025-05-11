package server

import "net/http"

func (s *server) IsAuthenticated(r *http.Request) bool {
	exists := s.session.Exists(r.Context(), "user_id")
	return exists
}

func (s *server) Auth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !s.IsAuthenticated(r) {
			s.logger.Printf("log in first!")
			// TODO: redirect or pass a statuscode and let htmx redirect
			return
		}
	})
}
