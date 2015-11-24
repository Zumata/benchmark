package main

import (
	"fmt"
	"testing"
)

func BenchmarkDeleteWhileLooping(b *testing.B) {

	for i := 0; i < 10000; i++ {
		deleteWhileLooping()

		if i%1000 == 0 {
			fmt.Print(".")
		}
	}

}

func BenchmarkDeleteByTrimming(b *testing.B) {

	for i := 0; i < 10000; i++ {
		deleteByTrimming()

		if i%1000 == 0 {
			fmt.Print(".")
		}
	}

}

func BenchmarkDeleteByAdding(b *testing.B) {

	for i := 0; i < 10000; i++ {
		deleteByAdding()

		if i%1000 == 0 {
			fmt.Print(".")
		}
	}

}
