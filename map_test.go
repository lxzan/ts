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

func TestMap_Delete(t *testing.T) {
	var c = NewMap()
	c.Set("name", "caster")
	c.Set("age", 12)
	c.Delete("name")
	println(c)
}

func TestMap_ForEach(t *testing.T) {
	var c = NewMap()
	c.Set("name", "caster")
	c.Set("age", 12)
	c.Set("sex", "male")
	c.Set("smoke", false)
	c.ForEach(func(k string, v interface{}) {
		print(k, v)
	})
}

func TestMap_Clear(t *testing.T) {
	var c = NewMap()
	c.Set("name", "caster")
	c.Set("age", 12)
	c.Set("sex", "male")
	c.Set("smoke", false)
	c.Clear()
	println(c)
}