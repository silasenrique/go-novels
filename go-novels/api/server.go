package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"

	"go-models/api/route"
)

type Server struct {
	server *chi.Mux
	db     *sql.DB
}

func NewServer(db *sql.DB) *Server {

	return &Server{
		server: chi.NewMux(),
		db:     db,
	}
}

func (s *Server) SetRoutes() *Server {
	s.server.Route("/", func(r chi.Router) {
		r.Mount("/site", route.SiteRoutes(s.db))
		// r.Route("/site", route.SiteRoutes(s.db))
	})

	return s
}

func (s *Server) Run(port string) {
	http.ListenAndServe(port, s.server)
}
