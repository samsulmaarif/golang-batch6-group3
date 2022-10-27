package model

type RajaOngkirDefault struct {
	RajaOngkir RajaOngkir `json:"rajaongkir"`
}

type RajaOngkir struct {
	Query   Query   `json:"query"`
	Status  Status  `json:"status"`
	Results Results `json:"results"`
}

type Query struct {
	Id          string `json:"id"`
	Province    string `json:"province"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	Courier     string `json:"courier"`
}

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Results struct {
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

// ///////////////// Costs /////////////////////

type RajaOngkirDefaultCost struct {
	RajaOngkir RajaOngkirCost `json:"rajaongkir"`
}

type RajaOngkirCost struct {
	Query              QueryCost          `json:"query"`
	Status             Status             `json:"status"`
	OriginDetails      OriginDetails      `json:"origin_details"`
	DestinationDetails DestinationDetails `json:"destination_details"`
	ResultsCost        []ResultsCost      `json:"results"`
}

type QueryCost struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	Courier     string `json:"courier"`
}

type OriginDetails struct {
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type DestinationDetails struct {
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type ResultsCost struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Costs []Costs `json:"costs"`
}

type Costs struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        []Cost `json:"cost"`
}

type Cost struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}
