package service

import (
	"golang-batch6-group3/adaptor"
	"golang-batch6-group3/server/params"
	"golang-batch6-group3/server/repository"
	"golang-batch6-group3/server/view"
)

type RajaOngkirServices struct {
	repo            repository.RajaOngkirRepo
	typicodeAdaptor *adaptor.TypicodeAdaptor
}

func NewRajaOngkirServices(repo repository.RajaOngkirRepo, typicodeAdaptor *adaptor.TypicodeAdaptor) *RajaOngkirServices {
	return &RajaOngkirServices{
		repo:            repo,
		typicodeAdaptor: typicodeAdaptor,
	}
}

func (ro *RajaOngkirServices) FindProvinceById(req *params.RajaOngkirQuery) *view.Response {
	query := req.ParseToModel()
	province, err := ro.repo.FindProvinceById(query)
	if err != nil {
		return view.ErrInternalServer("FIND_PROVINCE_FAIL", err.Error())
	}
	return view.SuccessFindAll(province)
}

func (ro *RajaOngkirServices) FindCityById(req *params.RajaOngkirQuery) *view.Response {
	query := req.ParseToModel()
	city, err := ro.repo.FindCityById(query)
	if err != nil {
		return view.ErrInternalServer("FIND_CITY_FAIL", err.Error())
	}
	return view.SuccessFindAll(city)
}

func (ro *RajaOngkirServices) FindCost(req *params.RajaOngkirQuery) *view.Response {
	query := req.ParseToModel()
	city, err := ro.repo.FindCost(query)
	if err != nil {
		return view.ErrInternalServer("FIND_COST_FAIL", err.Error())
	}
	return view.SuccessFindAll(city)
}
