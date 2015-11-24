package sliceindexes

import (
	"testing"
)

func BenchmarkLookupAndAssign(b *testing.B) {
	LookupAndAssign()
}

func BenchmarkCopyAndAssign(b *testing.B) {
	CopyAndAssign()
}
