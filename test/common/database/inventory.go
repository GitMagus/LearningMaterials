package database

import "github.com/beego/beego/v2/client/orm"

type Inventory struct {
	Id     int `orm:"auto"`
	Gid    int
	GCount int
	Status int8
}

func DescGoodCount(gid int, count int) bool {
	o := orm.NewOrm()
	var i Inventory
	if err := o.QueryTable("inventory").Filter("gid", gid).One(&i); err != nil {
		return false
	}
	if i.GCount < count {
		return false
	}
	i.GCount = i.GCount - count

	_, err := o.Update(&i)
	if err != nil {
		return false
	}
	return true

}
