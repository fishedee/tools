package decimal

import (
	"testing"

	"github.com/fishedee/tools/assert"
)

func TestDecimalOperation(t *testing.T) {
	testCase := []struct {
		handler func(Decimal, Decimal) Decimal
		left    string
		right   string
		result  string
	}{
		{Decimal.Add, "1.0000", "2.0000", "3"},
		{Decimal.Add, "1.1234", "2.1000", "3.2234"},
		{Decimal.Sub, "3.0000", "2.0000", "1"},
		{Decimal.Sub, "-3.0000", "2.0000", "-5"},
		{Decimal.Mul, "1.234", "5.6789", "7.0077626"},
		{Decimal.Mul, "-1.234", "5.6789", "-7.0077626"},
		{Decimal.Mul, "1.234", "-5.6789", "-7.0077626"},
		{Decimal.Mul, "-1.234", "-5.6789", "7.0077626"},
		{Decimal.Mul, "1.2330", "5.6789", "7.0020837"},
		{Decimal.Mul, "1.2320", "5.6789", "6.9964048"},
		{Decimal.Mul, "-1.2320", "5.6789", "-6.9964048"},
		{Decimal.Mul, "1.2320", "-5.6789", "-6.9964048"},
		{Decimal.Mul, "-1.2320", "-5.6789", "6.9964048"},
		{Decimal.Div, "1.0000", "2.0000", "0.5"},
		{Decimal.Div, "2.0000", "1.0000", "2"},
		{Decimal.Div, "1.0000", "7.0000", "0.1428571428571429"},
		{Decimal.Div, "-1.0000", "7.0000", "-0.1428571428571429"},
		{Decimal.Div, "1.0000", "-7.0000", "-0.1428571428571429"},
		{Decimal.Div, "-1.0000", "-7.0000", "0.1428571428571429"},
		{Decimal.Div, "1.0000", "3.0000", "0.3333333333333333"},
		{Decimal.Div, "-1.0000", "3.0000", "-0.3333333333333333"},
		{Decimal.Div, "1.0000", "-3.0000", "-0.3333333333333333"},
		{Decimal.Div, "-1.0000", "-3.0000", "0.3333333333333333"},
		{Decimal.Div, "1.0000", "9.0000", "0.1111111111111111"},
		{Decimal.Div, "0.0600", "78.0000", "0.0007692307692308"},
		{Decimal.Div, "0.0500", "78.0000", "0.0006410256410256"},
		{Decimal.Div, "0.0010", "78.0000", "0.0000128205128205"},
		{Decimal.Div, "0.0010", "-78.0000", "-0.0000128205128205"},
		{Decimal.Div, "-0.0010", "78.0000", "-0.0000128205128205"},
		{Decimal.Div, "-0.0010", "-78.0000", "0.0000128205128205"},
		//需要支持到千亿级别的大整数
		{Decimal.Mul, "1567891.2320", "356.6789", "559233719.9494048"},
		{Decimal.Mul, "-1567891.2320", "356.6789", "-559233719.9494048"},
		{Decimal.Mul, "1567891.2320", "-356.6789", "-559233719.9494048"},
		{Decimal.Mul, "-1567891.2320", "-356.6789", "559233719.9494048"},
		{Decimal.Add, "900719925474.0986", "0.0001", "900719925474.0987"},
		{Decimal.Sub, "900719925474.0987", "0.0001", "900719925474.0986"},
		{Decimal.Mul, "100079991719.3443", "9", "900719925474.0987"},
		{Decimal.Mul, "100079991719.3443", "9", "900719925474.0987"},
		{Decimal.Mul, "9", "100079991719.3443", "900719925474.0987"},
		{Decimal.Mul, "949062.6563", "949062.6563", "900719925583.21192969"},
		{Decimal.Div, "900719925474.0991", "100079991719.3443", "9.000000000000004"},
		{Decimal.Div, "900719925474.0991", "9", "100079991719.3443444444444444"},
		{Decimal.Div, "900719925473.1207", "949062.6563", "949062.656184000040504"},
	}
	for singleTestCaseIndex, singleTestCase := range testCase {
		leftDecimal, err := NewDecimal(singleTestCase.left)
		assert.Equal(t, err, nil)
		rightDecimal, err := NewDecimal(singleTestCase.right)
		assert.Equal(t, err, nil)
		resultDecimal, err := NewDecimal(singleTestCase.result)
		assert.Equal(t, err, nil)
		result := singleTestCase.handler(leftDecimal, rightDecimal)
		assert.Equal(t, result, resultDecimal, singleTestCaseIndex)
	}
}

func TestDecimalRound(t *testing.T) {
	testCase := []struct {
		origin    string
		precision int
		result    string
	}{
		{"2.0000", 0, "2"},
		{"2.1234", 0, "2"},
		{"2.4999", 0, "2"},
		{"-2.4999", 0, "-2"},
		{"2.5000", 0, "3"},
		{"-2.5000", 0, "-3"},
		{"2.5001", 0, "3"},
		{"3.2000", 1, "3.2"},
		{"3.2123", 1, "3.2"},
		{"3.2499", 1, "3.2"},
		{"-3.2499", 1, "-3.2"},
		{"3.2500", 1, "3.3"},
		{"-3.2500", 1, "-3.3"},
		{"3.2501", 1, "3.3"},
		{"4.3200", 2, "4.32"},
		{"4.3212", 2, "4.32"},
		{"4.3249", 2, "4.32"},
		{"-4.3249", 2, "-4.32"},
		{"4.325", 2, "4.33"},
		{"-4.3250", 2, "-4.33"},
		{"4.3251", 2, "4.33"},
		{"5.4320", 3, "5.432"},
		{"5.4321", 3, "5.432"},
		{"5.4324", 3, "5.432"},
		{"-5.4324", 3, "-5.432"},
		{"5.4325", 3, "5.433"},
		{"-5.4325", 3, "-5.433"},
		{"5.4325", 3, "5.433"},
	}

	for singleTestCaseIndex, singleTestCase := range testCase {
		originDecimal, err := NewDecimal(singleTestCase.origin)
		assert.Equal(t, err, nil)
		resultDecimal, err := NewDecimal(singleTestCase.result)
		assert.Equal(t, err, nil)
		result := originDecimal.Round(singleTestCase.precision)
		assert.Equal(t, result, resultDecimal, singleTestCaseIndex)
	}
}

func TestDecimalOther(t *testing.T) {
	//cmp
	assert.Equal(t, Decimal("10").Cmp(Decimal("12")), -1)
	assert.Equal(t, Decimal("12").Cmp(Decimal("12")), 0)
	assert.Equal(t, Decimal("15").Cmp(Decimal("12")), 1)

	//equal
	assert.Equal(t, Decimal("12").Equal(Decimal("12")), true)
	assert.Equal(t, Decimal("-12").Equal(Decimal("-12")), true)
	assert.Equal(t, Decimal("0").Equal(Decimal("")), true)
	assert.Equal(t, Decimal("").Equal(Decimal("0")), true)
	assert.Equal(t, Decimal("-0").Equal(Decimal("0")), true)
	assert.Equal(t, Decimal("-1").Equal(Decimal("0")), false)

	//Sign
	assert.Equal(t, Decimal("-12").Sign(), -1)
	assert.Equal(t, Decimal("0").Sign(), 0)
	assert.Equal(t, Decimal("12").Sign(), 1)

	//Abs
	assert.Equal(t, Decimal("-12").Abs(), Decimal("12"))
	assert.Equal(t, Decimal("0").Abs(), Decimal("0"))
	assert.Equal(t, Decimal("12").Abs(), Decimal("12"))
}

func TestDecimalToString(t *testing.T) {
	testCase := []struct {
		origin string
		result string
	}{
		{"", "0"},
		{"1", "1"},
		{"7.0078", "7.0078"},
		{"-7.0078", "-7.0078"},
		{"0.1", "0.1"},
		{"-0.1", "-0.1"},
		{"0.01", "0.01"},
		{"-0.01", "-0.01"},
		{"0.001", "0.001"},
		{"-0.001", "-0.001"},
		{"0.0001", "0.0001"},
		{"-0.0001", "-0.0001"},
		{"1.2345", "1.2345"},
		{"1.234549", "1.234549"},
		{"1.23455", "1.23455"},
		{"0.23459", "0.23459"},
		{"78559233719.9494", "78559233719.9494"},
		{"-78559233719.9494", "-78559233719.9494"},
	}
	for _, singleTestCase := range testCase {
		orginDecimal := Decimal(singleTestCase.origin)
		result := orginDecimal.String()
		assert.Equal(t, result, singleTestCase.result)
	}
}
