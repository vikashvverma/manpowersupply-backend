package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/vikashvverma/manpowersupply-backend/job"
	"github.com/vikashvverma/manpowersupply-backend/response"
)

func Upsert(f job.Fetcher, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var j job.Job
		json.NewDecoder(r.Body).Decode(&j)
		if validateJob(&j) {
			response.ClientError(w)
			return
		}

		rowsAffected, err := f.SaveAndUpdate(&j)
		if err != nil {
			l.WithError(err).Errorf("upsert: could not save/update job")
			response.ServerError(w)
			return
		}
		response.Success{Success: fmt.Sprintf("%d row saved successfully!", rowsAffected)}.Send(w)
	}
}

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

func DeleteJob(f job.Fetcher, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		jobIDString, ok := vars["jobID"]
		var jobID int
		var err error
		if ok {
			jobID, err = strconv.Atoi(jobIDString)
			if err != nil {
				l.WithError(err).Errorf("DeleteJob: invalid jobID supplied: %s", jobIDString)
				response.ClientError(w)
				return
			}
		}
		err = f.Delete(int64(jobID))
		if err != nil {
			l.WithError(err).Errorf("Delete: could not delete job: %d", jobID)
			response.ServerError(w)
			return
		}
		response.Success{Success: fmt.Sprintf("Job having jobID: %d deleted successfully", jobID)}.Send(w)
	}
}

func validateJob(j *job.Job) bool {
	if j.JobID <= 0 || len(j.Location) == 0 || len(j.Industry) == 0 || len(j.Title) == 0 || j.TypeID == 0 {
		return true
	}
	return false
}
