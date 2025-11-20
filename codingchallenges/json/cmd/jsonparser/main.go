package main

import (
	"os"

	"github.com/et-codes/lab/codingchallenges/json/internal/jsonparser"
)

func main() {
	jp := jsonparser.NewJSONParser(os.Args[1:])
	os.Exit(jp.Run())
}
