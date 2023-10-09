package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type GoodsType struct {
	Id   int    `orm:"column(id);auto" description:"商品分类ID"`
	Name string `orm:"column(name);size(20)" description:"类型"`
}

func (t *GoodsType) TableName() string {
	return "goods_type"
}

func init() {
	orm.RegisterModel(new(GoodsType))
}
