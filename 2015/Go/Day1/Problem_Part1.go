package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	leftHeader  = "#=================="
	rightHeader = "==================#"
	DEBUG       = false // Don't clutter up our output unless we are debugging
)

// '(' == up one floor
// ')' == down one floor

func main() {
	fmt.Printf("%s Welcome to Advent Of Code! Day 1, Part1 %s\n", leftHeader, rightHeader)
	parseArgs()
	logger := setupLogging("[INFO]: ")
	content, err := os.ReadFile("input.txt")
	if err != nil {
		logger.Fatal("Unable to read file: ", err)
		return
	}

	floor := 0
	/*
		 I don't think we even need a char conversion since, we can operate on the runes/literal
		inputText := string(content)
	*/
	for _, char := range content {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
		if DEBUG {
			fmt.Printf("Current Char: %v, Current Floor: %v\n", char, floor)
		}
	}
	logger.Printf("Final Floor - %v\n", floor)

}

func setupLogging(logLevel string) log.Logger {
	logger := log.New(os.Stdout, logLevel, log.Ldate|log.Ltime|log.Lmsgprefix)
	return *logger
}

func parseArgs() {
	flag.Parse()
}
