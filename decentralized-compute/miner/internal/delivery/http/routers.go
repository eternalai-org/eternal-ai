package http

import (
	"solo/internal/delivery/http/response"
)

func (h *httpDelivery) registerRoutes() {
	h.router.Use(response.LogRequest)
	api := h.router.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/chains", h.chains).Methods("GET")
	api.HandleFunc("/chains/{chain_id}/inferences", h.createInference).Methods("POST")
	api.HandleFunc("/health-check", h.healthCheck).Methods("GET")
	api.HandleFunc("/devices/information", h.information).Methods("GET")
	api.HandleFunc("/on-chain/data", h.onChainData).Methods("GET")
	api.HandleFunc("/nodes", h.listNode).Methods("GET")
	api.HandleFunc("/nodes", h.createNode).Methods("POST")
	api.HandleFunc("/nodes/{node_id}/inferences", h.listInference).Methods("GET")
	h.printRoutes()
}
