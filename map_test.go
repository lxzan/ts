package ts

import "testing"

func TestMap_Set(t *testing.T) {
	var c = NewMap()
	c.Set("name", "caster")
	println(c)
}

func TestMap_Get(t *testing.T) {
	var c = NewMap()
	c.Set("name", "caster")
	name, _ := c.Get("name")
	println(name)
}
