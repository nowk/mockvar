package mockvar

import (
	"testing"

	"gopkg.in/nowk/assert.v2"
)

func TestMockAndRestoreBasicTypes(t *testing.T) {
	{
		v := "One"
		r := Mock(&v, "Two")
		assert.Equal(t, "Two", v)

		r()
		assert.Equal(t, "One", v, "restore")
	}

	{
		v := 1
		r := Mock(&v, 2)
		assert.Equal(t, 2, v)

		r()
		assert.Equal(t, 1, v, "restore")
	}
}

func TestMockAndRestoreFunc(t *testing.T) {
	{
		v := func() string { return "One" }
		r := Mock(&v, func() string { return "Two" })
		assert.Equal(t, "Two", v())

		r()
		assert.Equal(t, "One", v(), "restore")
	}

	{
		v := func(n int) int { return n + 1 }
		r := Mock(&v, func(n int) int { return n + 2 })
		assert.Equal(t, 3, v(1))

		r()
		assert.Equal(t, 2, v(1), "restore")
	}
}
