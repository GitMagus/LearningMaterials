package pay_service

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"log"
	"months/common/models"
	"strconv"
)

var (
	appId        = "appid"
	privateKey   = "privateKey"
	aliPublicKey = "aliPublicKey"
)
var client *alipay.Client
var err error
var Client *alipay.Client

func init() {
	client, err = alipay.New(appId, privateKey, false)
	if err != nil {
		fmt.Println(err)
	}
	client.LoadAliPayPublicKey(aliPublicKey)
	Client = client
	log.Println(client)
}

func CreatePayUrl(o models.Order) string {
	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://127.0.0.1:8888/api/v1/notify"
	p.ReturnURL = "http://127.0.0.1:8888/api/v1/notify"
	p.Subject = o.Subject
	p.OutTradeNo = o.OrderNo
	p.TotalAmount = strconv.Itoa(int(float64(o.Count) * o.Fee))
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	log.Println(o)
	log.Println(p)
	var url, err = client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	fmt.Println(payURL)

	return payURL
}
