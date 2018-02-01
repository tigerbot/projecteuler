package misc

import (
	"encoding/csv"
	"os"
	"strconv"
)

// ReadNumberGrid reads base10 numbers from a file
func ReadNumberGrid(name string, sep rune) [][]int64 {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = sep

	raw, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	result := make([][]int64, len(raw))
	for i, line := range raw {
		parsedLine := make([]int64, len(line))
		for j, rawNum := range line {
			if num, err := strconv.ParseInt(rawNum, 10, 64); err != nil {
				panic(err)
			} else {
				parsedLine[j] = num
			}
		}
		result[i] = parsedLine
	}

	return result
}

// ReadTextGrid reads a grid of arbitrary strings from a file
func ReadTextGrid(name string, sep rune) [][]string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = sep

	raw, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return raw
}

// ReadFlattenText reads a csv and returns the flattened results
func ReadFlattenText(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	all, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	if len(all) == 1 {
		return all[0]
	}
	size := 0
	for _, line := range all {
		size += len(line)
	}
	result := make([]string, 0, size)
	for _, line := range all {
		result = append(result, line...)
	}
	return result
}
