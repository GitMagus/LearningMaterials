package main

import (
	"2102Amonth/common/database"
	escl "2102Amonth/common/elasticsearch"
)

func main() {
	// 同步mysql订单数据到es中
	ExportToEs()
}

func ExportToEs() {
	orderList := database.GetOrderList()
	for _, val := range orderList {
		escl.InsertData(val)
	}
}
