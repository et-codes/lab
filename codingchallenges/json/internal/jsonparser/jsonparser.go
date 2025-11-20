package jsonparser

import "fmt"

type JSONParser struct{}

func NewJSONParser() *JSONParser {
	return &JSONParser{}
}

func (j *JSONParser) Run() {
	fmt.Println("running application...")
}
