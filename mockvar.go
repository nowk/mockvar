package mockvar

import (
	"reflect"
)

type restoreFunc func()

// Mock sets b to a and provides restore func to set the original a back to a
// TODO maybe add some better panic management for less cryptic panic messages
// when things like unsettables are found, etc...
func Mock(a, b interface{}) restoreFunc {
	v := reflect.ValueOf(a)
	e := v.Elem()

	// create backup of original
	o := reflect.New(e.Type())
	o.Elem().Set(e)

	e.Set(reflect.ValueOf(b)) // set new value

	return func() {
		e.Set(o.Elem())
	}
}
