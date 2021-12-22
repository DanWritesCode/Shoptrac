package endpoints

import (
	"../data"
	"../response"
	"net/http"
)

func GetRevenue(w http.ResponseWriter, r *http.Request) {
	s := data.Revenue{
		Sales:           1587.12,
		ShippingCharged: 487.96,
		TaxesCollected:  115.74,
		Tips:            10,
	}
	response.JSON(w, 200, s)
}
