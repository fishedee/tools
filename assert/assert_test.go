package assert

import "testing"

func TestEqual(t *testing.T) {
	Equal(t, 0, 0)
	Equal(t, 1.0, 1.0)
	Equal(t, "a", "a")
}
