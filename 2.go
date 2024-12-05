package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func FindSafeReportsFromFile() int {
	// Open the file containing the reports
	f, err := os.Open("2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer f.Close()

	// Initialize safe report count
	safeCount := 0

	// Read file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into integers
		strNumbers := strings.Fields(line)
		var numbers []int
		for _, str := range strNumbers {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			numbers = append(numbers, num)
		}

		// Check if the report is safe
		isSafe:=isSafeReport(numbers)
		if isSafe {
			safeCount++
		}
	}

	// Check for errors while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return safeCount
}

// Check if a report is safe based on the given rules
func isSafeReport(numbers []int) bool {
	isIncreasing := false
	isDecreasing := false
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]

		// Not safe if the absolute difference is greater than 3
		if (math.Abs(float64(diff)) > 3) || (diff == 0) {
			return false
		}

		if diff > 0 {
			isIncreasing = true
		} else if diff < 0 {
			isDecreasing = true
		}
	}

	// Not safe if both increasing and decreasing
	return (isIncreasing && !isDecreasing) || (!isIncreasing && isDecreasing)
}

func main() {
	startTime := time.Now()
	
	safeReports := FindSafeReportsFromFile()
	fmt.Printf("Total safe reports: %d\n", safeReports)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
