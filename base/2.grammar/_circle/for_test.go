package _circle

import "testing"

func BenchmarkGeneral(b *testing.B) {
	for i := 0; i < b.N; i++ {
		general()
	}
}

func BenchmarkOtherWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		otherWay()
	}
}
