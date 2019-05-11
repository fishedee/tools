package subtest

import (
	"github.com/donnol/tools/query"
)

// Address Address
type Address struct {
	AddressId int
	City      string
}

func logic() {
	query.Column([]Address{}, "City")
}
