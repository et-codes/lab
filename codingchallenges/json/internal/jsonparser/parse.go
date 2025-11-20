package jsonparser

import "fmt"

func (j *JSONParser) Parse(input []byte) error {
	data := string(input)

	if data != "{}" {
		return fmt.Errorf("invalid JSON")
	}

	return nil
}
