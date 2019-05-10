package decimal

import (
	"github.com/shopspring/decimal"
)

// Decimal 浮点
type Decimal string

// NewDecimal 新建
func NewDecimal(in string) (Decimal, error) {
	_, err := decimal.NewFromString(in)
	if err != nil {
		return "", err
	}
	return Decimal(in), nil
}

func getDecimal(a Decimal) decimal.Decimal {
	if string(a) == "" {
		return decimal.Decimal{}
	}
	r, err := decimal.NewFromString(string(a))
	if err != nil {
		panic(err)
	}
	return r
}

// Add 加
func (left Decimal) Add(right Decimal) Decimal {
	l := getDecimal(left)
	r := getDecimal(right)
	return Decimal(l.Add(r).String())
}

// Sub 减
func (left Decimal) Sub(right Decimal) Decimal {
	l := getDecimal(left)
	r := getDecimal(right)
	return Decimal(l.Sub(r).String())
}

// Mul 乘
func (left Decimal) Mul(right Decimal) Decimal {
	l := getDecimal(left)
	r := getDecimal(right)
	return Decimal(l.Mul(r).String())
}

// Div 除
func (left Decimal) Div(right Decimal) Decimal {
	l := getDecimal(left)
	r := getDecimal(right)
	return Decimal(l.Div(r).String())
}

// Round 四舍五入
func (left Decimal) Round(precision int) Decimal {
	l := getDecimal(left)
	return Decimal(l.Round(int32(precision)).String())
}

// Cmp 比较
func (left Decimal) Cmp(right Decimal) int {
	l := getDecimal(left)
	r := getDecimal(right)
	return l.Cmp(r)
}

// Equal 相等
func (left Decimal) Equal(right Decimal) bool {
	l := getDecimal(left)
	r := getDecimal(right)
	return l.Equal(r)
}

// Sign 符号
func (left Decimal) Sign() int {
	l := getDecimal(left)
	return l.Sign()
}

// Abs 绝对值
func (left Decimal) Abs() Decimal {
	l := getDecimal(left)
	return Decimal(l.Abs().String())
}

// String 字符串
func (left Decimal) String() string {
	l := getDecimal(left)
	return l.String()
}
