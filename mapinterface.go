//recursive solution
package main

import (
	"fmt"
	"encoding/json"
)

type input interface{}

type structure struct {
	kind string
	payload payload
}

type payload struct {
	text string
	number float64
	truth bool
}

var testinput string = `{
  "kind": "response",
  "payload": {
    "text": "words",
    "number": 315.72,
    "truth": true
  }
}`

func Checker(input map[string]interface{}) {
	for _, v := range input {
		switch v.(type) {
		case map[string]interface{}:
			Checker(v.(map[string]interface{}))
		}
	}
}

func main() {
	var input input
	err := json.Unmarshal([]byte(testinput), &input)
	if err != nil {
		fmt.Println(err)
	}

	inputmap := input.(map[string]interface{})
	fmt.Println(inputmap)
	fmt.Printf("%T\n", inputmap)
	fmt.Println(inputmap["payload"])
	fmt.Printf("%T\n", inputmap["payload"])
	payload := inputmap["payload"]
	fmt.Println(payload)
	fmt.Printf("%T\n", payload)

	payloadValue := payload.(map[string]interface{})
	fmt.Println(payloadValue["text"])

	structure := input.(structure)
	fmt.Println(structure)
}
