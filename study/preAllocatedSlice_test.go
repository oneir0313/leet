package study

import "testing"

var aBillion = 1000 * 1000 * 1000

// Add items to a slice with a billion-item array allocated up-front

func BenchmarkBillionPreallocatedSlice(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		items := make([]uint8, 0, aBillion)
		for j := 0; j < aBillion; j++ {
			items = append(items, 255)
		}
	}
}

// Add items to a slice that starts out nil; the underlying array gets
// re-allocated and re-copied several times
func BenchmarkBillionEmptySlice(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var items []uint8
		for j := 0; j < aBillion; j++ {
			items = append(items, 255)
		}
	}
}
