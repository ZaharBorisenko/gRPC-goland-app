package handlers

import (
	api "github.com/ZaharBorisenko/z-gateway/gRPC/gen/go"
	"github.com/ZaharBorisenko/z-gateway/lib"
	"net/http"
)

type Handler struct {
	client api.OrderServiceClient
}

func NewHandler(client api.OrderServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customersID}/orders", h.HandleCreateOrder)
}

func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("customerID")

	var items []api.ItemsWithQuantity
	if err := lib.ReadJSON(r, &items); err != nil {
		lib.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	ptrItems := make([]*api.ItemsWithQuantity, len(items))
	for i := range items {
		ptrItems[i] = &items[i]
	}

	order, err := h.client.CreateOrder(r.Context(), &api.CreateOrderRequest{
		CustomerID: id,
		Items:      ptrItems,
	})
	if err != nil {
		lib.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	lib.WriteJSON(w, http.StatusCreated, order)
}
