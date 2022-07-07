package study

import (
	"reflect"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	res := sliceAppend()
	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 7, 8}) {
		t.Error("sliceAllocate is not correct")
	}
}
