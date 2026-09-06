package main

import "fmt"

/* class A {
	public ALiID string
	public Pay(money: number)(string, float64){
		xxxxxx
	}
} */

// 开头大写 代表公共的 所有的包都可以访问
// 开头小写 代表私有的
// 适用于所有的 变量 函数 结构体 声明
// interface 定义函数
// struct 定义属性

// 定义函数 interface
type Payment interface {
	Pay(value float64) (string, float64)
}

// 定义属性 struct
type ALiPay struct {
	ALiID string
}

func (a ALiPay) Pay(money float64) (string, float64) {
	return a.ALiID, money
}

type WeChatPay struct {
	WXID string
}

func (w WeChatPay) Pay(money float64) (string, float64) {
	return w.WXID, money
}

/**
 * 业务层封装

 * pay 函数结构体
 * value 金额
 */
func CreateOrder(pay Payment, value float64) {
	id, money := pay.Pay(value)
	fmt.Println(id, money)
}

func main4() {
	// 初始化完成 结构体 后 才可以正常使用
	/* ali := ALiPay{ALiID: "ali_123"}
	aliId, aliMoney := ali.Pay(666)
	fmt.Println(aliId, aliMoney) */

	CreateOrder(ALiPay{ALiID: "ali_123"}, 666)

	/* wx := WeChatPay{WXID: "wx_123"}
	wxId, wxMoney := wx.Pay(888)
	fmt.Println(wxId, wxMoney) */

	CreateOrder(WeChatPay{WXID: "wx_123"}, 888)
}
