package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"../db/query"
	"../utils"
	"github.com/go-chi/chi"
)

// GetProductPricesByBuyerID returns a list of prices of products
func GetProductPricesByBuyerID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "buyerId")
	ctx := context.Background()

	pricesProducts := query.SearchProductPriceByBuyerID(&ctx, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pricesProducts)
}

//GetSimilarPricesByTopLow return a list of products with similar prices
func GetSimilarPricesByTopLow(w http.ResponseWriter, r *http.Request) {
	avg := r.URL.Query().Get("avg")
	ctx := context.Background()

	// A variance of 10 percent
	top, low := utils.GetTopAndLowStr(avg, 8)

	similarProducts := query.SearchProductsByTopLow(&ctx, top, low)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(similarProducts)
}
