package models

import "github.com/google/uuid"

type Engine struct {
	ID             uuid.UUID `json:"engine_id"`
	Displacement   int64     `json:"displacement"`
	NoOfCyclinders int64     `json:"no_of_cyclinders"`
	CarRange       int64     `json:"car_range"`
}
