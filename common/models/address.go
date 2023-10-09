package models

type Address struct {
	Id       int    `orm:"column(id);auto" description:"地址ID"`
	Location string `orm:"column(location);size(100)" description:"地址、地点、位置"`
}

func (t *Address) TableName() string {
	return "address"
}
