package car

import (
	"CMS/models"
	"CMS/store"
	service "CMS/store/service"
	"context"
)

type CarService struct {
	store store.CarStoreInterface
}

// Ensure CarService implements service.CarServiceInterface
var _ service.CarServiceInterface = (*CarService)(nil)

func NewCarService(store store.CarStoreInterface) *CarService {
	return &CarService{
		store: store,
	}
}

func (s *CarService) GetCarByID(ctx context.Context, id string) (models.Car, error) {
	return s.store.GetCarById(ctx, id)
}

func (s *CarService) GetCarsByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	cars, err := s.store.GetCarByBrand(ctx, brand, isEngine)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (s *CarService) CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {
	if err := models.ValidateRequest(*carReq); err != nil {
		return models.Car{}, err
	}

	car, err := s.store.CreateCar(ctx, carReq)
	if err != nil {
		return models.Car{}, err
	}

	return car, nil
}

func (s *CarService) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error) {
	if err := models.ValidateRequest(*carReq); err != nil {
		return models.Car{}, err
	}

	car, err := s.store.UpdateCar(ctx, id, carReq)
	if err != nil {
		return models.Car{}, err
	}

	return car, nil
}

func (s *CarService) DeleteCar(ctx context.Context, id string) (models.Car, error) {
	car, err := s.store.DeleteCar(ctx, id)
	if err != nil {
		return models.Car{}, err
	}
	return car, nil
}
