package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	leftHeader  = "#=================="
	rightHeader = "==================#"
	DEBUG       = true // Don't clutter up our output unless we are debugging
)

func main() {
	fmt.Printf("\n\n%s Welcome to Advent Of Code 2023! Day 1, Part 2 %s\n\n", leftHeader, rightHeader)
	// Let's load our input file
	content, error := os.Open("input.txt")
	//content, error := os.Open("test.txt")
	if error != nil {
		fmt.Print("Unable to read file.: ", error)
	}
	defer content.Close() // Don't leave file descriptors open

	// Create a slice/array of ints; For each line's int we'll add that. So if line 1 has only one digit, 7, we'll add 77
	integerLineSlice := []int{}

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		line := scanner.Text()
		if DEBUG {
			log.Printf("Line: %s\n", line)
		}

		count := 1
		first_int := 000 // Since we are operating on single char's, we'll never "set" to this number in an edge case
		last_int := first_int
		fmt.Println("String: ", line)
		for _, char := range line {

			fmt.Printf("Working with char: %v. Count is: %v\n", string(char), count)
			// unicode.IsDigit also includes other numerals besides ASCII. Let's just make it easy and check if rune between 0 and 9
			if count == 1 && char >= '0' && char <= '9' {
				first_int = int(char)
				fmt.Println("this is the FIRST digit: ", string(first_int))
				count++
			} else if count > 1 && char >= '0' && char <= '9' {
				last_int = int(char)
				fmt.Println("this is a digit: ", string(last_int))
				count++
			}
		}
		// If last_int is still set to 000, then we assume the line only had one int, and thus set the last to the first.
		if last_int == 000 {
			last_int = first_int
		}
		fmt.Printf("First int: %v, Last int: %v\n", string(first_int), string(last_int))
		combinedStr := string(first_int) + string(last_int)
		combinedInt, _ := strconv.Atoi(combinedStr)
		fmt.Println("String: ", combinedStr)
		fmt.Printf("Combined Final Int: %v\n\n", combinedInt)
		integerLineSlice = append(integerLineSlice, combinedInt)
	}

	runningSum := 0
	for _, val := range integerLineSlice {
		runningSum += val
		fmt.Printf("Running Sum is: %v\n ", runningSum)
	}
	fmt.Println("Final Tally is: ", runningSum)

}
