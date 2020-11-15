package router

import (
	"github.com/gorilla/mux"
	"github.com/lulumel0n/m3u8-relay/server/middleware"
)

// Router for routing
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/music.m3u", middleware.GetStreaming).Methods("GET", "OPTIONS")

	return router
}
