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
	filenameArg := flag.Arg(0)
	if filenameArg == "" {
		fmt.Println("No filename provided")
		os.Exit(1)
	}

	filename, err := getPath(filenameArg)
	if err != nil {
		fmt.Println("Error parsing path to file:", err.Error())
		os.Exit(1)
	}

	// check that file exists
	_, err = os.Stat(filename)
	if err != nil {
		fmt.Printf("Cannot stat file %q: %s\n", filename, err.Error())
		os.Exit(1)
	}

	// process flags
	if *cFlag {
		n, err := countBytes(filename)
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%8d %s\n", n, filename)
	}

	if *lFlag {
		n, err := countLines(filename)
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%8d %s\n", n, filename)
	}
}

// countBytes returns the size of the file
func countBytes(filename string) (int, error) {
	path, err := getPath(filename)
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
func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
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
