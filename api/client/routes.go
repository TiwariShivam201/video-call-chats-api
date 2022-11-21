package client

import "github.com/gorilla/mux"

func LoadClientRoutes(router *mux.Router) {
	agoraTokenRoutes := router.PathPrefix("/client").Subrouter()

	//agora token generate
	agoraTokenRoutes.HandleFunc("/agora-token", GenerateAgoraToken).Methods("GET")

}
