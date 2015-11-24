package main

import (
	"bytes"
	"encoding/json"
)

type CustomMap1 map[string]interface{}
type CustomMap2 map[string]interface{}

func (a CustomMap1) MarshalJSON() (jsonBytes []byte, err error) {
	buf := new(bytes.Buffer)

	buf.WriteByte(byte('['))

	index := 0
	mapLen := len(a)
	for _, val := range a {
		var itemJSON []byte
		if itemJSON, err = json.Marshal(val); err != nil {
			return
		}
		buf.Write(itemJSON)

		index++
		if index < mapLen {
			buf.WriteByte(byte(','))
		}
	}

	buf.WriteByte(byte(']'))

	jsonBytes = buf.Bytes()
	return
}

func (a CustomMap2) MarshalJSON() (jsonBytes []byte, err error) {
	tempSlice := make([]interface{}, len(a))

	index := 0
	for key, _ := range a {
		tempSlice[index] = a[key]
		index++
	}

	jsonBytes, err = json.Marshal(tempSlice)
	return
}
