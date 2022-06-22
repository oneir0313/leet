package study

import (
	"reflect"
	"testing"
)

func TestSliceAllocate(t *testing.T) {
	res := sliceAllocate()
	if !reflect.DeepEqual(res, []int{1,2,3,4,5,7,8}) {
		t.Error("sliceAllocate is not correct")
	}
}
