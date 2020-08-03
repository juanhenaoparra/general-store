package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"../db/query"
	"github.com/go-chi/chi"
)

// BuyerCtx handle buyer route
// func BuyerCtx (next http.Handler) {
// 	return
// }

// GetBuyersByPage returns a list of buyers
func GetBuyersByPage(w http.ResponseWriter, r *http.Request) {
	first := r.URL.Query().Get("first")
	offset := r.URL.Query().Get("offset")
	ctx := context.Background()

	buyersGetted := query.GetBuyersPaginated(&ctx, first, offset)

	json.NewEncoder(w).Encode(buyersGetted)
}

// GetBuyerProfile returns a buyer profile with all data
func GetBuyerProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "buyerId")
	first := r.URL.Query().Get("first")
	offset := r.URL.Query().Get("offset")
	ctx := context.Background()

	buyerProfile := query.GetBuyerProfile(&ctx, id, first, offset)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buyerProfile)
}
