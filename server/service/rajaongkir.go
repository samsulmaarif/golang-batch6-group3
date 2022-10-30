package service

import (
	"golang-batch6-group3/server/params"
	"golang-batch6-group3/server/repository"
	"golang-batch6-group3/server/view"
)

type RajaOngkirServices struct {
	repo repository.RajaOngkirRepo
}

func NewRajaOngkirServices(repo repository.RajaOngkirRepo) *RajaOngkirServices {
	return &RajaOngkirServices{
		repo: repo,
	}
}

func (ro *RajaOngkirServices) FindProvinceById(req *params.RajaOngkirQueryProvince) *view.Response {
	query := req.ParseToModel()
	province, err := ro.repo.FindProvinceById(query)
	if err != nil {
		return view.ErrInternalServer("FIND_PROVINCE_FAIL", err.Error())
	}
	return view.SuccessFindAll("FIND_PROVINCE_SUCCESS", province)
}

func (ro *RajaOngkirServices) FindCityById(req *params.RajaOngkirQueryCity) *view.Response {
	query := req.ParseToModel()
	city, err := ro.repo.FindCityById(query)
	if err != nil {
		return view.ErrInternalServer("FIND_CITY_FAIL", err.Error())
	}
	return view.SuccessFindAll("FIND_CITY_SUCCESS", city)
}

func (ro *RajaOngkirServices) FindCost(req *params.RajaOngkirQueryCost) *view.Response {
	query := req.ParseToModel()
	city, err := ro.repo.FindCost(query)
	if err != nil {
		return view.ErrInternalServer("FIND_COST_FAIL", err.Error())
	}
	return view.SuccessFindAll("FIND_COST_SUCCESS", city)
}
