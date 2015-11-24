package main

import (
	"encoding/json"
	"strconv"
	"testing"
)

var result []byte
var data map[string]interface{}
var data1map5items CustomMap1
var data2map5items CustomMap2
var data1map1000items CustomMap1
var data2map1000items CustomMap2
var data1map1000000items CustomMap1
var data2map1000000items CustomMap2

func generateSampleMap(size int) map[string]interface{} {

	data := make(map[string]interface{}, size)

	for n := 0; n < size; n++ {
		data["what-"+strconv.Itoa(n)] = map[string]string{"yo": "yo"}
	}

	return data
}

func init() {

	data1map5items = CustomMap1(generateSampleMap(5))
	data2map5items = CustomMap2(generateSampleMap(5))

	data1map1000items = CustomMap1(generateSampleMap(1000))
	data2map1000items = CustomMap2(generateSampleMap(1000))

	data1map1000000items = CustomMap1(generateSampleMap(1000000))
	data2map1000000items = CustomMap2(generateSampleMap(1000000))

}

func BenchmarkCustomMapOneLoop5Items10(b *testing.B) {

	var jsonOut []byte
	var err error

	for n := 0; n < b.N; n++ {
		jsonOut, err = json.Marshal(data1map5items)
		if err != nil {
			panic(err)
		}
	}
	result = jsonOut

}

func BenchmarkCustomMapTwoLoop5Items10(b *testing.B) {

	var jsonOut []byte
	var err error

	for n := 0; n < b.N; n++ {
		jsonOut, err = json.Marshal(data2map5items)
		if err != nil {
			panic(err)
		}
	}
	result = jsonOut

}

func BenchmarkCustomMapOneLoop1000Items10(b *testing.B) {

	var jsonOut []byte
	var err error

	for n := 0; n < b.N; n++ {
		jsonOut, err = json.Marshal(data1map1000items)
		if err != nil {
			panic(err)
		}
	}

	result = jsonOut

}

func BenchmarkCustomMapTwoLoop1000Items10(b *testing.B) {

	var jsonOut []byte
	var err error

	for n := 0; n < b.N; n++ {
		jsonOut, err = json.Marshal(data2map1000items)
		if err != nil {
			panic(err)
		}
	}

	result = jsonOut

}

func BenchmarkCustomMapOneLoop1000000Items10(b *testing.B) {

	var jsonOut []byte
	var err error

	for n := 0; n < b.N; n++ {
		jsonOut, err = json.Marshal(data1map1000000items)
		if err != nil {
			panic(err)
		}
	}

	result = jsonOut

}

func BenchmarkCustomMapTwoLoop1000000Items10(b *testing.B) {

	var jsonOut []byte
	var err error

	for n := 0; n < b.N; n++ {
		jsonOut, err = json.Marshal(data2map1000000items)
		if err != nil {
			panic(err)
		}
	}

	result = jsonOut

}
