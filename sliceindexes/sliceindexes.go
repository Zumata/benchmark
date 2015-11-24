package sliceindexes

const (
	OuterCount int = 10000
	InnerCount int = 10000
)

type SomeStruct struct {
	Id     int
	Nested []NestedStruct
}

type NestedStruct struct {
	Count float32
}

func LookupAndAssign() {
	generatedValues := generator()

	var countSlice []int
	for _, individualSomeStruct := range generatedValues {
		for _, individualNestedStruct := range individualSomeStruct.Nested {
			countSlice = append(countSlice) * 1.01
		}
	}
}

func CopyAndAssign() {
	generatedValues := generator()
	for outerIndex, _ := range generatedValues {
		for innerIndex, _ := range generatedValues[outerIndex].Nested {
			_ = generatedValues[outerIndex].Nested[innerIndex].Count * 1.01
		}
	}
}

func generator() (generatedValue []SomeStruct) {
	for i := 0; i < OuterCount; i++ {
		var nestedValues []NestedStruct
		for j := 0; j < InnerCount; j++ {
			nestedValues = append(nestedValues, NestedStruct{
				Count: float32(j),
			})
		}
		generatedValue = append(generatedValue, SomeStruct{
			Id:     i,
			Nested: nestedValues,
		})
	}
	return
}
