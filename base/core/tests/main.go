package tests

import "reflect"

//IsEqual compares two elements to equal
func IsEqual(expected, actual interface{}) bool {
	return reflect.DeepEqual(expected, actual)
}
