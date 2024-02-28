package restaurant

import "github.com/CesarDelgadoM/api-reports/internal/api/branch"

type RestaurantData struct {
	UserId      uint   `json:"-"`
	Name        string `json:"name"`
	Founder     string `json:"founder"`
	Location    string `json:"location"`
	Country     string `json:"country"`
	Fundation   string `json:"fundation"`
	Headquarter string `json:"headquarter"`
}

type Restaurant struct {
	UserId      uint            `json:"-"`
	Name        string          `json:"name"`
	Founder     string          `json:"founder"`
	Location    string          `json:"location"`
	Country     string          `json:"country"`
	Fundation   string          `json:"fundation"`
	Headquarter string          `json:"headquarter"`
	Branches    []branch.Branch `json:"branches"`
}

func (rest *Restaurant) MapToRestaurantData() RestaurantData {
	return RestaurantData{
		Name:        rest.Name,
		Founder:     rest.Founder,
		Location:    rest.Fundation,
		Country:     rest.Country,
		Fundation:   rest.Fundation,
		Headquarter: rest.Headquarter,
	}
}

type Request struct {
	UserId     uint       `json:"userid"`
	Name       string     `json:"name"`
	Restaurant Restaurant `json:"restaurant"`
}
