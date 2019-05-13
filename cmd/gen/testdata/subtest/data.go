package subtest

import (
	"github.com/fishedee/tools/query"
)

// Address Address
type Address struct {
	AddressId int
	City      string
}

func logic() {
	query.Column([]Address{}, "City")
}
