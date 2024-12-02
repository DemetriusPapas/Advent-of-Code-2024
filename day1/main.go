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

func calculateSimilarity (loc1, loc2 []int) (totalSimilarity int) {

	// individual similarity of every element
	iSimilarity := make([]int, len(loc1))

	// calculate similarity for first element separately
	for j := range loc2 {
		if loc1[0] == loc2[j] {
			iSimilarity[0] ++
		}
	}
	totalSimilarity += loc1[0] * iSimilarity[0]

	// calculate for the rest of the elements.
	// since the list is ordered, check if an element is the same as
	// the previous one to avoid recalculating.
	for i := 1; i < len(loc1); i++ {
		if loc1[i] == loc1[i-1] {
			iSimilarity[i] = iSimilarity[i-1]
		} else {
			for j := range loc2 {
				if loc1[i] == loc2[j] {
					iSimilarity[i] ++
				}
			}
		}

		totalSimilarity += loc1[i] * iSimilarity[i]
	}

	return totalSimilarity
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

	filename := "input"
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

	// Part 1

	fmt.Printf("total distance: %v\n", tallyDistance(loc1, loc2))

	// Part 2

	fmt.Printf("total similarity: %v\n", calculateSimilarity(loc1, loc2))

}
