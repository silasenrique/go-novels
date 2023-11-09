package route

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"go-models/model"
	"go-models/service"
)

func SiteRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", getSites(db))
	router.Post("/", insertSite(db))

	return router
}

func getSites(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "minha mina ta irritada comigo!"}`))
	}
}

func insertSite(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		site := &model.Site{}
		err := json.NewDecoder(r.Body).Decode(site)
		if err != nil {
			errText := fmt.Sprintf(
				`{"error": "nao foi possivel ler o conteudo enviado. err %s"}`,
				err,
			)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errText))
			return
		}

		insertService := service.NewSiteInsertService(db)
		site, err = insertService.Insert(site)
		if err != nil {
			errText := fmt.Sprintf(`{"error": "%s"}`, err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errText))
			return
		}

		err = json.NewEncoder(w).Encode(site)
		if err != nil {
			errText := fmt.Sprintf(`{"error": "nao foi possivel formatar a reposta, err: %s"}`, err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errText))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
