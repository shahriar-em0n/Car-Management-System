package models

import (
	"errors"
	"strconv"
	"time"

	// "github.com/docker/docker/daemon/names"
	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fuel_type"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct{
	Name string `json:"name"`
	Year string `json:"year"`
	Brand string `json:"brand"`
	FuelType string `json:"fuel_type"`
	Engine Engine `json:"engine"`
	Price float64 `json:"price"`
}

func validateName(name string) error{
	if name == "" {
		return errors.New("Name is Required")
	}
	return nil
}

func validateYear(year string) error{
	if year == ""{
		return errors.New("Year is Required")
	}
	_, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("Year must be a valid number")
	}
	currentYear := time.Now().Year()
	yearInt, _:= strconv.Atoi(year)

	if yearInt < 1886 || yearInt > currentYear {
		return errors.New("Year must be between 1886 and current year")
	}
	return nil
}