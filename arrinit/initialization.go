package arrinit

const(
    SampleString = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func AppendStringArrayWithCapZero(appendCount int) {
    strArr := make([]string, 0, 0)

    for i := 0; i < appendCount; i++ {
        strArr = append(strArr, SampleString)
    }
}

func AppendStringArrayWithAppendCountCap(appendCount int) {
    strArr := make([]string, 0, appendCount)

    for i := 0; i < appendCount; i++ {
        strArr = append(strArr, SampleString)
    }
}

func PopulateStringArrayWithAssignCountLenCap(assignCount int) {
    strArr := make([]string, assignCount, assignCount)

    for i := 0; i < assignCount; i++ {
        strArr[i] = SampleString
    }
}
