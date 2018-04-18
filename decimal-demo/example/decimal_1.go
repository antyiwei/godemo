package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Decimal struct {
	// contains filtered or unexported fields
}

func main() {
	n, err := decimal.NewFromString("-123.4567")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("金额：%v", n.String()))
	fmt.Println("------------------------------------")
	Test1()
	fmt.Println("------------------------------------")

}

func Test1() {
	d1 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3))
	d1.String() // output: "0.6666666666666667"
	fmt.Println(fmt.Sprintf("金额：%v", d1))

	d2 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(30000))
	d2.String() // output: "0.0000666666666667"
	fmt.Println(fmt.Sprintf("金额：%v", d2))

	d3 := decimal.NewFromFloat(20000).Div(decimal.NewFromFloat(3))
	d3.String() // output: "6666.6666666666666667"
	fmt.Println(fmt.Sprintf("金额：%v", d3))

	// 尝试小数的位数
	decimal.DivisionPrecision = 3
	d4 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3))
	d4.String() // output: "0.667"
	fmt.Println(fmt.Sprintf("金额：%v", d4))

	// 取负
	d5 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3)).Neg()
	d5.String() // output: "0.667"
	fmt.Println(fmt.Sprintf("金额：%v", d5))

	// 将点往后移动两位
	d6 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3)).Shift(2)
	d6.String() // output: "0.667"
	fmt.Println(fmt.Sprintf("金额：%v", d6))

	// 平均值
	d7 := decimal.Avg(decimal.New(12, -2), decimal.New(23, -1))
	d7.String()
	fmt.Println(fmt.Sprintf("平均值金额：%v", d7))
}
