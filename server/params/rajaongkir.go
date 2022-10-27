package params

import (
	"golang-batch6-group3/server/model"
)

type RajaOngkirQuery struct {
	ProvinceId  string
	CityId      string
	Origin      string
	Destination string
	Weight      int
	Courier     string
}

func (rq *RajaOngkirQuery) ParseToModel() *model.Query {
	return &model.Query{
		Id:          rq.CityId,
		Province:    rq.ProvinceId,
		Origin:      rq.CityId,
		Destination: rq.Destination,
		Weight:      rq.Weight,
		Courier:     rq.Courier,
	}
}
