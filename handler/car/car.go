package car

import "CMS/service"


type CarHandler struct{
	service service.CarServiceInterface
}

func NewCarHandler(service service.CarServiceInterface) *CarHandler{
	return &CarHandler{
		service: service,
	}
}