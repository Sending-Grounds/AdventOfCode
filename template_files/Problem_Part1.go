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
	fmt.Printf("%s Welcome to Advent Of Code! Day 1, Part 2 %s\n", leftHeader, rightHeader)
	parseArgs()
	logger := setupLogging("[INFO]: ")
	content, err := os.ReadFile("Input.txt")
	if err != nil {
		logger.Fatal("Unable to read file: ", err)
		return
	}

	logger.Printf("Content:\n%s\n", content)
}

func setupLogging(logLevel string) log.Logger {
	logger := log.New(os.Stdout, logLevel, log.Ldate|log.Ltime|log.Lmsgprefix)
	return *logger
}

func parseArgs() {
	flag.Parse()
}
