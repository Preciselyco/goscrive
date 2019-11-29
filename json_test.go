package scrive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseJsonDocument(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/document.json")
	if err != nil {
		t.Fail()
	}
	d := &Document{}
	if err := parseJson(b, d); err != nil {
		fmt.Printf("parseJson error: %s\n", err)
		t.Fail()
	}
	marshaled, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal: %s\n", err)
	}
	fmt.Printf("%s\n", marshaled)
}
