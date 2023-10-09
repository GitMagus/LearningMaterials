package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Inventory struct {
	Gcount int  `orm:"column(gcount)" description:"商品数量"`
	Gid    int  `orm:"column(gid)" description:"商品ID"`
	Id     int  `orm:"column(id);auto" description:"库存ID"`
	Status int8 `orm:"column(status)" description:"库存状态"`
}

func (t *Inventory) TableName() string {
	return "inventory"
}

func init() {
	orm.RegisterModel(new(Inventory))
}
