package planetsvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetPlanetEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetPlanetEndpoint: MakeGetPlanetEndpoint(s),
	}
}

func MakeGetPlanetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getPlanetRequest)
		p, e := s.GetPlanet(ctx, req.s)
		return getPlanetResponse{s: p, Err: e}, nil
	}
}

type getPlanetRequest struct {
	s string
}

type getPlanetResponse struct {
	s   string `json:"s"`
	Err error  `json:"err,omitempty"`
}

func (r getPlanetResponse) error() error { return r.Err }
