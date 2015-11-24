package main

import (
	"fmt"
	// "log"
	"strings"
)

func deleteWhileLooping() {
	testContent := generateSlice()

	originalLength := len(testContent)
	for tokenIndex, _ := range testContent {
		if tokenIndex%2 == 0 {

			if tokenIndex == len(testContent)-1 {
				testContent = testContent[:tokenIndex-(originalLength-len(testContent))-1]
			} else {
				testContent = append(testContent[:tokenIndex-(originalLength-len(testContent))], testContent[tokenIndex-(originalLength-len(testContent))+1:]...)
			}
		}
	}

	// log.Printf("While looping(%v): %v", len(testContent), testContent)
}

func deleteByTrimming() {
	testContent := generateSlice()
	for tokenIndex, _ := range testContent {
		if tokenIndex%2 == 0 {
			testContent[tokenIndex] = ""
		}
	}

	sparseResult := strings.Join(testContent, " ")
	strings.Fields(sparseResult)

	// log.Printf("While trimmed(%v): %v", len(trimmedSlice), trimmedSlice)
}

func deleteByAdding() {
	testContent := generateSlice()

	addedTokens := make([]string, 100)
	for tokenIndex, _ := range testContent {
		if tokenIndex%2 != 0 {
			addedTokens = append(addedTokens, testContent[tokenIndex])
		}
	}

	// log.Printf("While trimmed(%v): %v", len(addedTokens), addedTokens)
}

func generateSlice() []string {
	var contentSlice []string
	for i := 0; i < 200; i++ {
		contentSlice = append(contentSlice, fmt.Sprintf("%v", i))
	}
	return contentSlice
}
