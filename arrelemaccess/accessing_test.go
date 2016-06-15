package arrelemaccess_test

import "testing"

const (
    SampleCount int = 40000000
)

type SampleObj struct {
    Age int
}

var sampleObjs []SampleObj = initSampleArray()

func BenchmarkCopyElemAccess(b *testing.B) {
    for _, obj := range sampleObjs {
        _ = obj
    }
}

func BenchmarkElemIndexAccess(b *testing.B) {
    for i, _ := range sampleObjs {
        _ = sampleObjs[i]
    }
}

func initSampleArray() []SampleObj {
    strArr := make([]SampleObj, SampleCount, SampleCount)

    for i := 0; i < SampleCount; i++ {
        strArr[i] = SampleObj{
            Age: 22,
        }
    }

    return strArr
}
