package gorm_postgres

import (
	"encoding/json"
	"golang-batch6-group3/server/model"
	"golang-batch6-group3/server/repository"
	"net/http"

	"gorm.io/gorm"
)

type rajaOngkirRepo struct {
	db *gorm.DB
}

func NewRajaOngkirRepoGormPostgres(db *gorm.DB) repository.RajaOngkirRepo {
	return &rajaOngkirRepo{
		db: db,
	}
}

var rootUrl = "https://api.rajaongkir.com/starter/"
var key = "8ed7415f889eb79a249d1e38c6759377"

func (ro *rajaOngkirRepo) FindProvinceById(query *model.Query) (*model.RajaOngkirDefault, error) {
	url := rootUrl + "province?id=" + query.Province
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("key", key)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rajaongkir := model.RajaOngkirDefault{}
	err = json.NewDecoder(resp.Body).Decode(&rajaongkir)
	if err != nil {
		return nil, err
	}
	return &rajaongkir, nil
}

func (ro *rajaOngkirRepo) FindCityById(query *model.Query) (*model.RajaOngkirDefault, error) {
	url := rootUrl + "city?id=" + query.Id + "&province=" + query.Province
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("key", key)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rajaongkir := model.RajaOngkirDefault{}
	err = json.NewDecoder(resp.Body).Decode(&rajaongkir)
	if err != nil {
		return nil, err
	}
	return &rajaongkir, nil
}

// belom
func (ro *rajaOngkirRepo) FindCostById(query *model.Query) (*model.RajaOngkirDefault, error) {
	url := rootUrl + "city?id=" + query.Id + "&province=" + query.Province
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("key", key)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rajaongkir := model.RajaOngkirDefault{}
	err = json.NewDecoder(resp.Body).Decode(&rajaongkir)
	if err != nil {
		panic(err)
	}
	return &rajaongkir, nil
}
