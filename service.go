package planetsvc

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type Service interface {
	GetPlanet(ctx context.Context, p string) (string, error)
}

type PlanetService struct{}

func NewPlanetService() Service {
	return &PlanetService{}
}

func (ps *PlanetService) GetPlanet(ctx context.Context, p string) (string, error) {
	fmt.Print(" hello ")
	return "", nil
}
