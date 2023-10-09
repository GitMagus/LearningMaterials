package pay_service

import (
	"2102Amonth/common/database"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	_ "github.com/smartwalle/alipay/v3"
	"log"
	"strconv"
)

var (
	appId        = "9021000127662405"
	privateKey   = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCoYn/EdfSLDetaAQxgZEWJiyEDnNTYQyWTELINFVl3PcVy3IDJRUK+oFICXb4vu4u5fJqu5STykky+5zXB2RJiMECgrH8Ewbnk0lv7GGP5AzpFMxaUcWYfWvmN1Hwx9xd5lE85oWZgHVX1HD+d/OF2tRYo65xInb0DS8lTkQ2XdYYuCaobXrYIt/630v6sc8zsd5AOm3DVBkzeqN1GjPn7ZwAnUdgyHnJmP86URshvocC+59y7d1uoVppWXvXHrKxGnkRf/9Zg2X7Uznlcw6bSRcOTlmWRt2pcnHNo1aZ5Q4UDJnwLGX3JFMC7efAOGBCtsVYJ3DeJ70l0kZjKb8v/AgMBAAECggEAFcpRYHonQEt31njmQh8XxJnlltMX1Di0RxjVAc5uenja8WdJn0JARez3T+kTsJejmEhguBtdMarvFZFUKpbNwzkdU+2erl3nkY4Lay4X2eSqCMR4upMGNS6PtuP1gn3IYupnajzRw8RoGEUkkmxkPPloOblryxXmnOOU1SSjTqMoEdDGrYyHUiqmoLBbmCQ5kN/HVmWArlitTFl+yPN5b3f/HsS+Xp7LTLVNDayd5fXvU9Fn8jQNhW6qGcmSYK3kSHrfsALBU0YgEYA9tn30iU8Ag5Mpph+aEDj+DPZLYL0LGmD5avSwcUdvVfUZtKAq7VtZN+5CCAYUmeYx1978IQKBgQDRyi4WoU/tiV8qHJHL0mTteiv5zoTMuH7UjZVF0iSdqSm2zEQNab1FLrHak27/GJO9wmoLmYyW82bHv2L/VHw+N5ti1/YYaoGoosKumDt+H+dS4r7A2LhMJzJf/u9RKDRHvUX3pPUXuh+97JRbOyaC2YPAcetSgdY3gOU1wpAfbwKBgQDNeYkI+kckioQuup4R2Clxz6W+VLWvIjFz/9rHdzy7QVKcEetgfWzM4ZG0CrF221x0P5IjfUHcW3Zm5+Pmcbtu+7AjdY0obJ//vdLycGLCPZEBcPNBjk6mcHwmAlL9S7v2oIVGEZ7gu1Fb32UX7R4uXh2hFetXfC/AHt/cONbUcQKBgQCOn+BjfDnFqM3s6E1qf1gkZxm8lF2OfKOlR7hDJpHEKfx1DTv6zndhsFQz4GXmwcO3j5Pe3P5KpaGrjP71zW9GMTHPqjjh8o7Boh2u60oO/gubOxIqM2xgQYmq37u9thKM7y3BJgGGpv/rdAqMV5NgBBbhX9F3X1POmi+6M9MduwKBgA0iMJUnUjbOt8Y43XOsinGg4So6Reas29ZbS0OmpnYdpcceChp/yd2KqYKFkHRVbyJrEc886WHJYtcPCW0oZd+hLNAHan9j+hIhswMWFenuf39FFfdhvjs7Q2Q2j9Fq7tfDyOECrVIWBwxFSjPuTxiNoX0zZniQEolIL+zmoSyxAoGBAIPgS+hu9Ryqp25xfbg2XSjvXEyACfz9dEuwTu4YQMUPnvMU4Tk9B3is/rFScjF6F4iVLlOhsHX5PaQzky9ZFnzR+0zS8CVnrywwF/whoVUXeUCC31+vT/x6G6kEQxfFR1ewWlQiaEMJjSR5Udj611ojvLckF0+c0PYZPuK5yAHL"
	aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxMYpKn9sKgsOdcOCcSk0cIrtltURy/n4eESwPfxaHi98kJ65F2a1ZqHe339S2Wca0oUr4mu455RTby6lyAzbtsc2eH20UsrOf0caTDuFv16QNPYz8Qx6MSfqcC95/SSbvalT5/iw0KEZPa6bdxU9uaVl+Gu/JvqtJqa0K0Hno1oYKVrK2NPM9lRrDnzjL1KDwM1ffe2YbSerxTqsuXdq2FFuwzkR1fIHtclZIS1xHGX7HFNiWh38l11LtQFG828VV0JPnk7zppJyh66RJnuVhkdYjkSZ+hECIjpOmV2ftozBnqc7n4cfFULvxjwvNwc7Om6sxlcg3d8fMED0xUb6TQIDAQAB"
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

func CreatePayUrl(o database.Order) string {
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
