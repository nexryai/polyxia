package details

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/nexryai/polyxia/server/controller"
	"github.com/nexryai/polyxia/server/converter"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	debug := false
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryParams := r.URL.Query().Get("id")
	if queryParams == "" {
		w.WriteHeader(400)
		return
	}

	debugFlag := r.URL.Query().Get("debug")
	if debugFlag == "dummy" {
		debug = true
	}

	data, err := controller.GetEventDetailsJson(queryParams, debug)
	if err != nil || data == nil {
		if errors.Is(err, converter.ErrSlightSeaLevelChangeIsNotSupported) {
			// 若干の海面変動は緊急性が低いため
			w.WriteHeader(204)
			return
		}

		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", *data)
}