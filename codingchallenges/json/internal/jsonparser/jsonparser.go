package jsonparser

import (
	"fmt"
	"os"
)

type JSONParser struct {
	args []string
}

func NewJSONParser(args []string) *JSONParser {
	return &JSONParser{
		args: args,
	}
}

func (j *JSONParser) Run() int {
	if err := j.validateArgs(); err != nil {
		fmt.Println(err)
		return 1
	}

	reader, err := os.Open(j.args[0])
	if err != nil {
		fmt.Println(err)
		return 1
	}
	defer reader.Close()
	return 0
}

func (j *JSONParser) validateArgs() error {
	if len(j.args) == 0 {
		return fmt.Errorf("must provide path to JSON file")
	}
	return nil
}
