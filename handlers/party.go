package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/vikashvverma/manpowersupply-backend/party"
	"github.com/vikashvverma/manpowersupply-backend/response"
)

func SaveParty(f party.Fetcher, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p party.Party
		json.NewDecoder(r.Body).Decode(&p)
		if validateParty(&p) {
			response.ClientError(w)
			return
		}

		err := f.Save(&p)
		if err != nil {
			l.WithError(err).Errorf("SaveParty: could not save party")
			response.ServerError(w)
			return
		}
		response.Success{Success: "saved successfully!"}.Send(w)
	}
}
func validateParty(p *party.Party) bool {
	if len(p.Name) == 0 || len(p.Mobile) < 10 {
		return true
	}
	return false
}
func FindParty(f party.Fetcher, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		page := 0
		var err error
		pageString, ok := vars["page"]
		if ok {
			page, err = strconv.Atoi(pageString)
			if err != nil {
				response.ServerError(w)
				return
			}
			if page < 0 {
				response.ClientError(w)
				return
			}
		}

		parties, err := f.Fetch(int64(page))
		if err != nil {
			l.WithError(err).Errorf("findParty: could not fetch parties")
			response.ServerError(w)
			return
		}
		response.Success{parties}.Send(w)
	}
}
