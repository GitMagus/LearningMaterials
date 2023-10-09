package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"创建时间"`
	Id         int       `orm:"column(id);pk" description:"用户ID"`
	IdCard     string    `orm:"column(id_card);size(18)" description:"大陆公民身份证件号码"`
	Mobile     string    `orm:"column(mobile);size(11)" description:"大陆11位手机号"`
	Password   string    `orm:"column(password);size(64)" description:"hash加密64位密码"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime)" description:"更新时间"`
	Username   string    `orm:"column(username);size(30)" description:"用户名"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}
