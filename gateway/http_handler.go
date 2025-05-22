package main

import (
	"errors"
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

	if err := valiteItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	o, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		OrderID: customerID,
		Items:   items,
	})

	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	common.WriteJSON(w, http.StatusOK, o)
}

func valiteItems(items []*pb.ItemsWithQuantity) error {
	if len(items) == 0 {
		return common.ErrNoItems
	}

	for _, item := range items {
		if item.ID == "" {
			return errors.New("item ID cannot be empty")
		} 

		if item.Quantity <= 0 {
			return errors.New("item quantity must be greater than zero")
		} 
	}

	return nil
}