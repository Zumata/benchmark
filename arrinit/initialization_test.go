package arrinit_test

import (
    "testing"

    . "github.com/Zumata/benchmark/arrinit"
)

const (
    // SampleCount 10m
    SampleCount int = 10000000
)

func BenchmarkAppendStringArrayWithCapZero(b *testing.B) {
    AppendStringArrayWithCapZero(SampleCount)
}

func BenchmarkAppendStringArrayWithAppendCountCap(b *testing.B) {
    AppendStringArrayWithAppendCountCap(SampleCount)
}

func BenchmarkPopulateStringArrayWithAssignCountLenCap(b *testing.B) {
    PopulateStringArrayWithAssignCountLenCap(SampleCount)
}
