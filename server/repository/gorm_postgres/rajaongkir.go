package gorm_postgres

import (
	"encoding/json"
	"golang-batch6-group3/server/model"
	"golang-batch6-group3/server/repository"
	"net/http"
	"strconv"
	"strings"

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

func (ro *rajaOngkirRepo) FindCost(query *model.Query) (*model.RajaOngkirDefaultCost, error) {
	url := rootUrl + "cost"
	payloadUrl := "origin=" + query.Origin + "&destination=" + query.Destination + "&weight=" + strconv.Itoa(query.Weight) + "&courier=" + query.Courier
	payload := strings.NewReader(payloadUrl)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("key", key)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rajaongkir := model.RajaOngkirDefaultCost{}
	err = json.NewDecoder(resp.Body).Decode(&rajaongkir)
	if err != nil {
		return nil, err
	}
	return &rajaongkir, nil
}
