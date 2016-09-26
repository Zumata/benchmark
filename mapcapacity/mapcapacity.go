package mapcapacity

const (
	SampleString string = "test value"
)

func MapWithNoCap(testCap int) {
	strArr := make(map[int]string)

	for i := 0; i < testCap; i++ {
		strArr[i] = SampleString
	}
}

func MapWithCap(testCap int) {
	strArr := make(map[int]string, testCap)

	for i := 0; i < testCap; i++ {
		strArr[i] = SampleString
	}
}
