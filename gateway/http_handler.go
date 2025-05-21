package main

import (
	"net/http"

	"github.com/DavidReque/order-management-system/common"
	pb "github.com/DavidReque/order-management-system/common/api"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.handleCreateOrder)
}

func (h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items []*pb.ItemsWithQuantity

	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
	}

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		OrderID: customerID,
		Items:   items,
	})
}
