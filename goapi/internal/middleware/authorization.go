package middleware

import (
	"errors"
	"net/http"

	"github.com/emmajiugo/goapi/api"
	"github.com/emmajiugo/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		if database, err = tools.NewDatabase(); err != nil {
			log.Error("Failed to connect to database: ", err)
			api.InternalServerErrorHandler(w, err)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || token != (loginDetails).AuthToken) {
			log.Error("Unauthorized access attempt by user: ", username)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}