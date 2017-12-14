package router

import (
	"github.com/gorilla/mux"

	"github.com/vikashvverma/manpowersupply-backend/config"
	"github.com/vikashvverma/manpowersupply-backend/factory"
	"github.com/vikashvverma/manpowersupply-backend/handlers"
	"github.com/vikashvverma/manpowersupply-backend/healthcheck"
)

const (
	GET  = "GET"
	POST = "POST"
)

func Router(c *config.Config, f *factory.Factory) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck.Self).Methods(GET)
	router.HandleFunc("/api/manpower/party", handlers.SaveParty(f.Fetcher(), f.Logger())).Methods(POST)
	router.HandleFunc("/api/manpower/party", handlers.FindParty(f.Fetcher(), f.Logger())).Methods(GET)
	router.HandleFunc("/api/manpower/party/{page}", handlers.FindParty(f.Fetcher(), f.Logger())).Methods(GET)
	//router.HandleFunc("/api/techscan/{lang}", handler.Language(f.Fetcher(), f.Logger())).Methods(GET)
	//router.HandleFunc("/api/techscan/{lang}/{page}", handler.Language(f.Fetcher(), f.Logger())).Methods(GET)
	//router.HandleFunc("/api/owner/{repoID}", handler.Owner(f.Fetcher(), f.Logger())).Methods(GET)
	return router
}
