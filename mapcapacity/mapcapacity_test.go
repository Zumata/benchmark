package mapcapacity_test

import (
	"testing"

	. "github.com/Zumata/benchmark/mapcapacity"
)

const (
	TestCap int = 500000
	//TestCap int = 200000
	//TestCap int = 1000000
)

func BenchmarkMapWithNoCap(b *testing.B) {
	MapWithNoCap(TestCap)
}

func BenchmarkMapWithCap(b *testing.B) {
	MapWithCap(TestCap)
}
