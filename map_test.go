package ts

import "testing"

func TestMap_Set(t *testing.T) {
	var c = NewMap()
	c.Set("name", "caster")
	println(c)
}
