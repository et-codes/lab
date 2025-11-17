package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// define and parse flags
	cFlag := flag.Bool("c", false, "number of bytes in file")
	lFlag := flag.Bool("l", false, "number of lines in the file")
	flag.Parse()

	// get filename argument
	fileArg := flag.Arg(0)
	if fileArg == "" {
		errorAndExit(fmt.Errorf("no filename provided"))
	}

	filePath, err := getPath(fileArg)
	if err != nil {
		newErr := fmt.Errorf("error parsing path to file %q: %s", fileArg, err.Error())
		errorAndExit(newErr)
	}

	// check that file exists
	_, err = os.Stat(filePath)
	if err != nil {
		newErr := fmt.Errorf("cannot stat file %q: %s", filePath, err.Error())
		errorAndExit(newErr)
	}

	// process flags
	if *cFlag {
		n, err := countBytes(filePath)
		if err != nil {
			errorAndExit(err)
		}
		fmt.Printf("%8d %s\n", n, fileArg)
	}

	if *lFlag {
		n, err := countLines(filePath)
		if err != nil {
			errorAndExit(err)
		}
		fmt.Printf("%8d %s\n", n, fileArg)
	}
}

// countBytes returns the size of the file
func countBytes(filePath string) (int, error) {
	path, err := getPath(filePath)
	if err != nil {
		return 0, err
	}

	file, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	return int(file.Size()), nil
}

// countLines returns the number of lines in the file
func countLines(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	return lines, nil
}

// getPath returns the cleaned absolute path of the filename argument
func getPath(filename string) (string, error) {
	path, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}
	return path, nil
}

func errorAndExit(err error) {
	fmt.Println("ERROR:", err.Error())
	os.Exit(1)
}
