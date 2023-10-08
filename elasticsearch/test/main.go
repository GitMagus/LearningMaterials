package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"xg/goods/rpc/model/elasticsearch"
)

func main() {
	info, err := elasticsearch.SearchGoodsInfo("GoodsName", "iPhone", "goods")
	if err != nil {
		logs.Error(err)
		return
	}
	fmt.Println(info)
}
