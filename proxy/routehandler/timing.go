package routehandler

import (
	"net/http"
	"strconv"
)

// Handle StatsD Timing request
func (routeHandler *RouteHandler) handleTimingRequest(w http.ResponseWriter, r *http.Request, key string) {

	// get timing
	time, err := strconv.ParseInt(r.PostFormValue("time"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid time specified", 400)
	}

	// get sample rate
	var sampleRate float64 = 1
	sampleRatePostFormValue := r.PostFormValue("sampleRate")
	if sampleRatePostFormValue != "" {
		var err error
		sampleRate, err = strconv.ParseFloat(sampleRatePostFormValue, 32)
		if err != nil {
			http.Error(w, "Invalid sample rate specified", 400)
		}
	}

	// send request
	routeHandler.statsdClient.Timing(key, time, float32(sampleRate))
}
