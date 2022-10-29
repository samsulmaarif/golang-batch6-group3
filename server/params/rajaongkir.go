package params

import (
	"golang-batch6-group3/server/model"
)

type RajaOngkirQueryProvince struct {
	ProvinceId string
}

type RajaOngkirQueryCity struct {
	ProvinceId string
	CityId     string
}

type RajaOngkirQueryCost struct {
	Origin      string
	Destination string
	Weight      int
	Courier     string
}

func (rq *RajaOngkirQueryProvince) ParseToModel() *model.QueryProvince {
	return &model.QueryProvince{
		Id: rq.ProvinceId,
	}
}

func (rq *RajaOngkirQueryCity) ParseToModel() *model.QueryCity {
	return &model.QueryCity{
		Id:       rq.CityId,
		Province: rq.ProvinceId,
	}
}

func (rq *RajaOngkirQueryCost) ParseToModel() *model.QueryCost {
	return &model.QueryCost{
		Origin:      rq.Origin,
		Destination: rq.Destination,
		Weight:      rq.Weight,
		Courier:     rq.Courier,
	}
}
