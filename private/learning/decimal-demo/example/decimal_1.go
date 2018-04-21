package main

import (
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

func main() {
	n, err := decimal.NewFromString("-123.4567")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("金额：%v", n.String()))
	fmt.Println("------------------------------------")
	Test1()
	fmt.Println("------------------------------------")

	// 初始化
	InitDecimal()

	fmt.Println("------------------------------------")
	// 测试方法
	TestDecilimalMode()
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

// InitDecimal 初始化数据方法
func InitDecimal() {
	// New returns a new fixed-point decimal, value * 10 ^ exp.
	d1 := decimal.New(12345, -1)
	d1.String() // output: "1234.5"
	fmt.Println(fmt.Sprintf("New金额：%v", d1))

	// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
	d2 := decimal.NewFromBigInt(big.NewInt(12345), -2)
	d2.String() // output: "1234.5"
	fmt.Println(fmt.Sprintf("NewFromBigInt金额：%v", d2))

	// NewFromFloat converts a float64 to Decimal.
	//
	// Example:
	//
	//     NewFromFloat(123.45678901234567).String() // output: "123.4567890123456"
	//     NewFromFloat(.00000000000000001).String() // output: "0.00000000000000001"
	//
	// NOTE: some float64 numbers can take up about 300 bytes of memory in decimal representation.
	// Consider using NewFromFloatWithExponent if space is more important than precision.
	//
	// NOTE: this will panic on NaN, +/-inf
	d3 := decimal.NewFromFloat(1234.50)
	d3.String() // output: "1234.5"
	fmt.Println(fmt.Sprintf("NewFromBigInt金额：%v", d3))

	// NewFromString returns a new Decimal from a string representation.
	d4, err := decimal.NewFromString("-123.45")
	if err != nil {
		fmt.Println(fmt.Sprintf("NewFromBigInt err:：%v", err.Error()))
	}
	d4.String() // output: "-123.45"
	fmt.Println(fmt.Sprintf("NewFromBigInt金额：%v", d4))

	// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
	// number of fractional digits.
	//
	// Example:
	//
	//     NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
	//
	d5 := decimal.NewFromFloatWithExponent(123.436, -1)
	d5.String()
	fmt.Println(fmt.Sprintf("NewFromFloatWithExponent金额：%v", d5))

	// RequireFromString returns a new Decimal from a string representation
	// or panics if NewFromString would have returned an error.
	//
	// Example:
	//
	//     d := RequireFromString("-123.45")
	//     d2 := RequireFromString(".0001")
	//
	d6 := decimal.RequireFromString("-123.45")
	d6.String()
	fmt.Println(fmt.Sprintf("RequireFromString金额：%v", d6))
}

// TestDecilimalMode
func TestDecilimalMode() {

	// Abs returns the absolute value of the decimal.
	d6 := decimal.RequireFromString("-123.45")
	d6.Abs().String()
	fmt.Println(fmt.Sprintf("Abs金额：%v", d6))

	d71, err := decimal.NewFromString("123.45")
	if err != nil {
		fmt.Println(fmt.Sprintf("NewFromString err:：%v", err.Error()))
	}
	d72, err := decimal.NewFromString("123.45")
	if err != nil {
		fmt.Println(fmt.Sprintf("NewFromString err:：%v", err.Error()))
	}

	// Add returns d + d2.
	d80 := d71.Add(d72)
	d80.String()
	fmt.Println(fmt.Sprintf("Add金额：%v", d80))

	// Sub returns d - d2.
	d81 := d71.Sub(d72)
	d81.String()
	fmt.Println(fmt.Sprintf("Sub 金额：%v", d81))

	// Neg returns -d.
	d7, _ := decimal.NewFromString("123.45")
	d7.Neg()
	d7.String()
	fmt.Println(fmt.Sprintf("Neg 金额：%v", d7))

	// Mul returns d * d2.
	d82 := d71.Mul(d72)
	d82.String()
	fmt.Println(fmt.Sprintf("Mul 金额：%v", d82))

	// Shift
	d2 := decimal.NewFromFloat(123.45)
	d2.Shift(2).String()
	fmt.Println(fmt.Sprintf("Shift 金额：%v", d2))

	// Div returns d / d2. If it doesn't divide exactly, the result will have
	// DivisionPrecision digits after the decimal point.
	d3 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3))
	d3.String()
	fmt.Println(fmt.Sprintf("Div 金额：%v", d3))
}
