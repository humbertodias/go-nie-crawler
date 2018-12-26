package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteJSON(arr interface{}, filename string) {
	jsonMarshal, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	dst := &bytes.Buffer{}
	if err = json.Indent(dst, jsonMarshal, "", "\t"); err != nil {
		fmt.Fprintf(os.Stderr, "problem formatting: %s", err)
	}
	err = ioutil.WriteFile(filename, dst.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
