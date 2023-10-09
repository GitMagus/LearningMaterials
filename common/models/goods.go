package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Goods struct {
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"创建时间"`
	Describe   string    `orm:"column(describe);size(100)" description:"描述"`
	GoodsNo    string    `orm:"column(goods_no);size(50)" description:"商品编号"`
	Id         int       `orm:"column(id);auto" description:"商品ID"`
	Status     int8      `orm:"column(status)" description:"商品状态"`
	Title      string    `orm:"column(title);size(40)" description:"商品标题"`
	TypeId     int       `orm:"column(type_id)" description:"分类ID"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime)" description:"更新时间"`
}

func (t *Goods) TableName() string {
	return "goods"
}

func init() {
	orm.RegisterModel(new(Goods))
}
