package subtest

import (
	"github.com/fishedee/tools/query"
)

// Address Address
type Address struct {
	AddressID int
	City      string
}

func logic() {
	query.Column[Address, string]([]Address{}, "City")
}
