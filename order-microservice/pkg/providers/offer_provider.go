package providers

import (
	"net/http"
	"order-microservice/pkg/routes/external_router"
	"order-microservice/pkg/utils/http_error"
	"order-microservice/pkg/utils/logger"
)

type Offer struct {
	ID       uint    `json:"id"`
	SellerID uint    `json:"sellerId"`
	Gtin     int     `json:"gtin"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func NewOfferProvider() *Offer {
	return &Offer{}
}

func (offer Offer) Provide(route *external_router.ExternalRoute, token *string) *http_error.HttpError {
	err := FetchRequest(offer, *route, *token)

	if err != nil {
		logger.Log("could not fetch offer", err)
		return &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": "could not fetch offer"}}
	}

	if httpError := offer.ValidateProvided(); httpError != nil {
		return httpError
	}

	return nil
}

func (offer Offer) ValidateProvided() *http_error.HttpError {
	if offer.ID == 0 {
		logger.Log("offer id is zero", nil)
		return &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": "could not fetch offer"}}
	}

	return nil
}