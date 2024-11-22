package model
type Country struct{
	ID uint `gorm:"column:idx;primarykey"`
	Name string `gorm:"column:name;size:255"`
}
func (c *Country) TableName() string  {
	//real table in database
	return "country"
}