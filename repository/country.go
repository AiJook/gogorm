// crud country
package repository

import (
	"gogorm/model"

	"gorm.io/gorm"
)

// oop inheritance
// Interface ข้อบังคับว่าต้องมี function ตามนี้ => 	ต้อง override มาเขยีนให้เสร็จ
type CountryRepository interface {
	GetAll() (*[]model.Country, error)
	GetByName(name string) (*[]model.Country, error)
	GetByID(idx int) (*model.Country, error)
	DeleteByID(idx int) *gorm.DB
	InsertCountry(data model.Country) *gorm.DB
}

type countryDB struct {
	db *gorm.DB
}

func NewCountryRepository(gormdb *gorm.DB) CountryRepository {
	return countryDB{db: gormdb}
}
func (c countryDB) DeleteByID(idx int) *gorm.DB {
	country := model.Country{}
	result := c.db.Where("idx = ?", idx).Delete(&country)
	return result
}
func (c countryDB)InsertCountry(data model.Country) *gorm.DB{
	country := data
	result := c.db.Create((&country))
	return result

}
func (c countryDB) GetAll() (*[]model.Country, error) {
	countries := []model.Country{}
	result := c.db.Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	return &countries, nil

}
func (c countryDB) GetByName(name string) (*[]model.Country, error) {
	countries := []model.Country{}
	name = "%" + name + "%"
	result := c.db.Where("Name like  ?", name).Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	return &countries, nil
}
func (c countryDB) GetByID(idx int) (*model.Country, error) {
	country := model.Country{}
	result := c.db.Where("idx = ?", idx).Find(&country)
	if result.Error != nil {
		return nil, result.Error
	}
	return &country, nil
}
