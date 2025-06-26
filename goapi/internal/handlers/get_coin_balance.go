package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emmajiugo/goapi/api"
	"github.com/emmajiugo/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	if err = decoder.Decode(&params, r.URL.Query()); err != nil {
		log.Error("Failed to decode parameters: ", err)
		api.InternalServerErrorHandler(w, err)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase() 
	if err != nil {
		log.Error("Failed to connect to database: ", err)
		api.InternalServerErrorHandler(w, err)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error("No login details found for user: ", params.Username)
		api.InternalServerErrorHandler(w, nil)
		return
	}

	response := api.CoinBalanceResponse{
		Code:    http.StatusOK,
		Balance: (*tokenDetails).Coins,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error("Failed to encode response: ", err)
		api.InternalServerErrorHandler(w, err)
		return
	}
}