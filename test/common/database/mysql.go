package database

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"log"
)

func init() {
	//dsn := "root:root@tcp(127.0.0.1:3306)/2102A"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//log.Println(db)
	//if err != nil {
	//	fmt.Println("数据库链接错误")
	//}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/2102A")
	orm.RegisterModel(new(Order), new(Inventory))
	orm.RunSyncdb("default", false, true)

	orm.Debug = true
}

type Order struct {
	Id        int    `gorm:"primaryKey"`
	OrderNo   string `gorm:"column:order_no"`
	Uid       int
	Gid       int
	Count     int
	TotalFee  float64 `gorm:"column:total_fee"`
	Fee       float64
	PayType   int8   `gorm:"column:pay_type"`
	CreateAt  string `gorm:"column:create_at"`
	Status    int8
	AddressId int    `gorm:"column:address_id"`
	UpdatedAt string `gorm:"column:updated_at"`
	Subject   string
	TradeNo   string
}

//func (Order) Tablename() string {
//	return "order"
//}

func InsertOrder(o Order) {
	oo := orm.NewOrm()
	_, err := oo.Insert(&o)
	if err != nil {
		log.Println(err)
	}
}

func GetOrderInfo(orderNo string) Order {
	var oo Order
	o := orm.NewOrm()
	if err := o.QueryTable("order").Filter("order_no", orderNo).One(&oo); err != nil {
		return oo
	}
	return oo
}

func UpdateOrder(o Order) {
	oo := orm.NewOrm()
	_, err := oo.Update(&o)
	if err != nil {
		log.Println(err)
	}
}

func DeleteOrder() {
	o := orm.NewOrm()
	var oo Order
	if err := o.QueryTable("order").OrderBy("-id").One(&oo); err != nil {

	}
	o.Delete(&oo)
}

func GetOrderList() (orderList []Order) {
	o := orm.NewOrm()
	var oo []Order
	o.QueryTable("order").All(&oo)
	return oo
}

func GetRefundOrderList() (refundOrderList []Order) {
	o := orm.NewOrm()
	var oo []Order
	o.QueryTable("order").Filter("status", 2).All(&oo)
	return oo
}
