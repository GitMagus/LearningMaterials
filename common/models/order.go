package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Order struct {
	AddressId  int       `orm:"column(address_id)" description:"地址ID"`
	Count      int       `orm:"column(count)" description:"购买数量"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"订单创建时间"`
	Fee        float64   `orm:"column(fee);digits(10);decimals(2)" description:"单价"`
	Gid        int       `orm:"column(gid)" description:"商品ID"`
	Id         int       `orm:"column(id);auto" description:"订单主键ID"`
	OrderNo    string    `orm:"column(order_no);size(50)" description:"订单编号"`
	PayType    int8      `orm:"column(pay_type)" description:"支付状态0：待支付1：已支付2：支付失败3：待退款4：已退款"`
	Status     int8      `orm:"column(status)" description:"订单状态"`
	Subject    string    `orm:"column(subject);size(30)" description:"订单标题"`
	TotalFee   float64   `orm:"column(total_fee);digits(10);decimals(2)" description:"总费用"`
	TradeNo    string    `orm:"column(trade_no);size(30);null" description:"支付单号"`
	Uid        int       `orm:"column(uid)" description:"用户ID"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime)" description:"订单修改时间"`
}

func (t *Order) TableName() string {
	return "order"
}

func init() {
	orm.RegisterModel(new(Order))
}
