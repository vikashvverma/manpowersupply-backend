package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"

	"github.com/vikashvverma/manpowersupply-backend/party"
	"github.com/vikashvverma/manpowersupply-backend/response"
)

func FindAllParty(f party.Fetcher, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parties, err := f.Fetch()
		if err != nil {
			l.WithError(err).Errorf("findAllPart: could not fetch parties")
			response.ServerError(w)
			return
		}

		res, err := json.Marshal(parties)
		if err != nil {
			e := fmt.Sprintf("error marshalling response: %s", err)

			l.Error(e)
			response.ServerError(w)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(res)
	}
}
