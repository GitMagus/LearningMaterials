package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"strings"
	"xg/goods/rpc/model/mysql"
)
//类似goods替换需要的信息
func SearchGoodsInfo(key, val, index string) ([]mysql.Goods, error) {
	var data = map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{ //Todo：match_all查询全部  match查询指定字段
				key: val, //Todo：key查询字段  val查询的值
			},
		},
		//Todo：高亮 (highlight 高亮关键字)
		//"highlight": map[string]interface{}{
		//	"fields": map[string]interface{}{
		//		key: map[string]interface{}{},
		//	},
		//},
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res, err := EsConn.Search(
		EsConn.Search.WithContext(context.Background()),
		EsConn.Search.WithIndex(index),
		EsConn.Search.WithBody(strings.NewReader(string(marshal))),
		EsConn.Search.WithTrackTotalHits(true),
		EsConn.Search.WithPretty(),
		//Todo 分页
		//EsConn.Search.WithFrom(0), //首页从 0 开始
		//EsConn.Search.WithSize(2), //每页数量
	)
	if err != nil {
		logs.Error("ERROR: %s", err)
		return nil, err
	}
	//fmt.Println(res)

	//Todo：第一次取值
	var m = map[string]interface{}{}
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return nil, err
	}
	//fmt.Println("m")
	//fmt.Println(m)

	//Todo：第二次取值
	i := m["hits"].(map[string]interface{})["hits"].([]interface{})
	//fmt.Println("i")
	//fmt.Println(i)

	//Todo：第三次取值,定义结构体将取出的 map 数据序列化至结构体
	//Todo：定义对应的结构体切片，将其 es 数据 append 至切片内返回
	var g mysql.Goods
	var goods []mysql.Goods

	for _, v := range i {
		ms := v.(map[string]interface{})["_source"].(map[string]interface{})
		//Todo：ms的值：map[age:23 name:李四 phone:13233333332]
		//Todo：map 无法直接将值反序列化至结构体 先将其序列化为字节
		bytes, err := json.Marshal(ms)
		if err != nil {
			logs.Error("map序列化为字节失败")
			return nil, err
		}
		//Todo：在将其反序列化至结构体
		err = json.Unmarshal(bytes, &g)
		fmt.Println(g)
		if err != nil {
			logs.Error("将其反序列化至结构体失败")
			return nil, err
		}
		//Todo：将反序列化后的结构体追加切片返回用户
		goods = append(goods, g)
	}
	//延迟关闭
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	return goods, nil

}
