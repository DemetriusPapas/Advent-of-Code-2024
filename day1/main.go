package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Reads a file to the end and returns a slice of strings where
// each string corresponds to a single line from the file.
func readFileToSlice(filename string) ([]string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading from file: %v", err)
	}

	return lines, nil
}

// Goes through the input line by line and assigns each location ID to its corresponding
// slice, after converting it to an integer.
func extractLocationLists(lines []string, delimiter string) (loc1, loc2 []int, err error) {

	for _, line := range lines {

		s1, s2 := splitLine(line, delimiter)

		d1, err := strconv.Atoi(s1)
		if err != nil {
			return nil, nil, err
		}

		d2, err := strconv.Atoi(s2)
		if err != nil {
			return nil, nil, err
		}

		loc1 = append(loc1, d1)
		loc2 = append(loc2, d2)
	}

	return loc1, loc2, nil
}

// Takes two sorted slices as input and returns the total sum of
// distances between all the pairs of their corresponding elements.
func tallyDistance(loc1, loc2 []int) (totalDistance int) {

	for i := range loc1 {
		totalDistance += Abs(loc1[i] - loc2[i])
	}

	return totalDistance
}

func splitLine(line, delimiter string) (string, string) {

	split := strings.SplitN(line, delimiter, 2)
	return split[0], split[1]
}

func Abs(x int) int {

	if x < 0 {
		return -x
	}

	return x
}

func main() {

	filename := "input.in"
	delimiter := "   "

	// Reads each line of the input into a separate string in a slice.
	lines, err := readFileToSlice(filename)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	// Generates two unsorted slices, one for each list of location IDs.
	loc1, loc2, err := extractLocationLists(lines, delimiter)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	
	slices.Sort(loc1)
	slices.Sort(loc2)

	fmt.Printf("%v\n", tallyDistance(loc1, loc2))
}
