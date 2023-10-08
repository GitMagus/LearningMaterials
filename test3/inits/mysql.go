package inits

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", MConfig.Mysql.Username+":"+MConfig.Mysql.Password+"@tcp("+MConfig.Mysql.Host+":"+strconv.Itoa(MConfig.Mysql.Port)+")/"+MConfig.Mysql.Database)
	orm.RegisterModel(new(User), new(Integral))

	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

// 用户表
type User struct {
	Id       int    `orm:"auto"`
	Username string `orm:"size(32)"`
	Password string `orm:"size(32)"`
	Integral int
}

// 积分表
type Integral struct {
	Id       int `orm:"auto"`
	UserId   int
	Integral int
	Type     int8
}

func UpdateUserIntegral(userId int, score int) {
	var u User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("id", userId).One(&u)
	if u.Id <= 0 {
		// insert 新数据
		u.Integral = score
		o.Insert(&u)
	}
	u.Integral = score
	o.Update(&u)
}
