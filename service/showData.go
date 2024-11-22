package service

import (
	"gogorm/model"
	"gogorm/repository"

	"gorm.io/gorm"
)

type showDataService interface {
	PrintAllCountries() (*[]model.Country, error)
	GetAllCountriesByName(name string) (*[]model.Country, error)
	GetCountryByID(idx int) (*model.Country, error)
	DeleteByID(idx int) *gorm.DB
	InsertCountry(data model.Country) *gorm.DB
}
type showData struct {
	db *gorm.DB
}

func NewShowDataService(gormdb *gorm.DB) showDataService {
	return showData{db: gormdb}

}
func (s showData) InsertCountry(data model.Country) *gorm.DB {
	country := repository.NewCountryRepository(s.db)
	InsertData := country.InsertCountry(data)
	return InsertData
}
func (c showData) PrintAllCountries() (*[]model.Country, error) {
	countryRepo := repository.NewCountryRepository(c.db)
	countries, err := countryRepo.GetAll()
	if err != nil {
		return nil, err
	}
	// for _, country := range *countries {
	// 	fmt.Printf("%v", country)
	// }
	return countries, nil

}

func (c showData) GetAllCountriesByName(name string) (*[]model.Country, error) {
	countryRepo := repository.NewCountryRepository(c.db)
	countries, err := countryRepo.GetByName(name)
	if err != nil {
		return nil, err
	}
	// for _, country := range *countries {
	// 	fmt.Printf("%v", country)
	// }
	return countries, nil

}
func (c showData) GetCountryByID(idx int) (*model.Country, error) {
	countryRepo := repository.NewCountryRepository(c.db)
	country, err := countryRepo.GetByID(idx)
	if err != nil {
		return nil, err
	}
	return country, nil

}
func (c showData) DeleteByID(idx int) *gorm.DB {
	countryRepo := repository.NewCountryRepository(c.db)
	err := countryRepo.DeleteByID(idx)
	if err != nil {
		return err
	}
	return nil
}
