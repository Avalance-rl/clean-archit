package http_server

import "clean-artchit/internal/domain"

type API struct {
	serviceStore domain.ServiceStore
	serviceProd  domain.ServiceProduct
}

func NewAPI(
	serviceStore domain.ServiceStore,
	serviceProduct domain.ServiceProduct,
) *API {
	return &API{
		serviceStore: serviceStore,
		serviceProd:  serviceProduct,
	}
}
