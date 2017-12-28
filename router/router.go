package router

import (
	"github.com/gorilla/mux"

	"github.com/vikashvverma/manpowersupply-backend/config"
	"github.com/vikashvverma/manpowersupply-backend/factory"
	"github.com/vikashvverma/manpowersupply-backend/handlers"
	"github.com/vikashvverma/manpowersupply-backend/healthcheck"
)

const (
	DELETE = "DELETE"
	GET    = "GET"
	PATCH  = "PATCH"
	POST   = "POST"
	PUT    = "PUT"
)

func Router(c *config.Config, f *factory.Factory) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck.Self).Methods(GET)
	router.HandleFunc("/api/manpower/job", handlers.Upsert(f.JobFetcher(), f.Logger())).Methods(POST)
	router.HandleFunc("/api/manpower/job", handlers.Upsert(f.JobFetcher(), f.Logger())).Methods(PATCH)
	router.HandleFunc("/api/manpower/job", handlers.Upsert(f.JobFetcher(), f.Logger())).Methods(PUT)
	router.HandleFunc("/api/manpower/job/{jobID}", handlers.DeleteJob(f.JobFetcher(), f.Logger())).Methods(DELETE)
	router.HandleFunc("/api/manpower/job", handlers.FindJob(f.JobFetcher(), f.Logger())).Methods(GET)
	router.HandleFunc("/api/manpower/job", handlers.FindJob(f.JobFetcher(), f.Logger())).Queries("page", "{page}", "limit", "{limit}", "type", "{type}").Methods(GET)
	router.HandleFunc("/api/manpower/job/{id}", handlers.FindJob(f.JobFetcher(), f.Logger())).Methods(GET)
	router.HandleFunc("/api/manpower/job/industry/", handlers.Industry(f.JobFetcher(), f.Logger())).Methods(GET)
	router.HandleFunc("/api/manpower/job/type/", handlers.JobType(f.JobFetcher(), f.Logger())).Methods(GET)
	router.HandleFunc("/api/manpower/job/type/{id}", handlers.JobType(f.JobFetcher(), f.Logger())).Methods(GET)
	router.HandleFunc("/api/manpower/party", handlers.SaveParty(f.PartyFetcher(), f.Logger())).Methods(POST)
	router.HandleFunc("/api/manpower/party", handlers.FindParty(f.PartyFetcher(), f.Logger())).Methods(GET)
	router.HandleFunc("/api/manpower/party", handlers.FindParty(f.PartyFetcher(), f.Logger())).Queries("page", "{page}").Methods(GET)
	router.HandleFunc("/api/manpower/party/{id}", handlers.FindParty(f.PartyFetcher(), f.Logger())).Methods(GET)
	return router
}
