package api

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/gargath/mongoose/pkg/auth"
	"github.com/gargath/mongoose/pkg/backend"
)

func NewFromConfig(c *Config, s *sessions.CookieStore, b *backend.Backend) (*API, error) {
	api := &API{}
	api.prefix = c.Prefix
	api.b = b
	api.s = s
	api.sessionName = c.SessionName
	return api, nil
}

func (a *API) AddRoutes(router *mux.Router) {
	apiRouter := router.PathPrefix(a.prefix).Subrouter()
	apiRouter.Use(auth.TokenVerifierMiddleware)

	apiRouter.HandleFunc("/users", a.ListUsersHandler).Methods("GET")
	apiRouter.HandleFunc("/users", a.InsertUserHandler).Methods("POST")

}
