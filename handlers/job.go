package handlers

import (
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/vikashvverma/manpowersupply-backend/job"
	"github.com/vikashvverma/manpowersupply-backend/response"
)

func FindJob(f job.Fetcher, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		jobID := vars["id"]

		queryParams := r.URL.Query()

		page := 0
		limit := 10
		var err error
		pageString := queryParams["page"]
		limitString := queryParams["limit"]
		if len(jobID) == 0 && len(pageString) > 0 {
			page, err = strconv.Atoi(pageString[0])
			if err != nil {
				response.ServerError(w)
				return
			}
			if page < 0 {
				response.ClientError(w)
				return
			}
		}

		if len(jobID) == 0 && len(limitString) > 0 {
			limit, err = strconv.Atoi(limitString[0])
			if err != nil {
				response.ServerError(w)
				return
			}
			if limit < 0 {
				response.ClientError(w)
				return
			}
		}

		jobs, err := f.Fetch(jobID, int64(page), int64(limit))
		if err != nil {
			l.WithError(err).Errorf("findJob: could not fetch jobs")
			response.ServerError(w)
			return
		}
		response.Success{Success: jobs}.Send(w)
	}
}
