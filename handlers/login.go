package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/Sirupsen/logrus"

	"github.com/vikashvverma/manpowersupply-backend/config"
	"github.com/vikashvverma/manpowersupply-backend/response"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *config.Config, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)
		if c.Username() == u.Username && c.Password() == u.Password {
			response.Success{Success: true}.Send(w)
			return
		}

		response.ClientError(w)
	}
}
