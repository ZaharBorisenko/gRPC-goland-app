package handlers

import (
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customersID}/orders", h.HandleCreateOrder)
}

func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {

}
