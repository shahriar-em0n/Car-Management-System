package car

import (
	"CMS/models"
	"CMS/store"
	"context"
)

type CarService struct {
	store store.CarStoreInferface
}

func NewCarService(store store.CarStoreInferface) *CarService {
	return &CarService{
		store: store,
	}
}

func (s *CarService) GetByID(ctx context.Context, id string) (*models.Car, error) {
	car, err := s.store.GetCarById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (s *CarService) GetCarsByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	cars, err := s.store.GetCarByBrand(ctx, brand, isEngine)
	if err != nil {
		return nil, err
	}

	return cars, nil

}

func (s *CarService) CreateCar(ctx context.Context, car *models.CarRequest) (*models.Car, error) {
	if err := models.ValidateRequest(*car); err != nil {
		return nil, err
	}

	createdCar, err := s.store.CreateCar(ctx, car)
	if err != nil {
		return nil, err
	}
	return &createdCar, nil
}


func(s *CarService) UpdateCar(ctx, context.Context, id string, carReq *models.CarRequest)(*models.Car, error){
	if err := models.ValidateRequest(*carReq); err != nil {
		return nil, err
	}

	updatedCar, err := s.store.UpdateCar(ctx, id, *&carReq)
	if err != nil {
		return nil, err
	}

	return &updatedCar, nil

}