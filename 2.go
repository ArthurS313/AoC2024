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

func FindSafeReportsFromFile() (int, int) {
	// Open the file containing the reports
	f, err := os.Open("2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, 0
	}
	defer f.Close()

	// Initialize safe report counts
	safeCount := 0
	safeCount2 := 0

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

		// Check if the report is safe with the first set of requirements
		isSafe:=isSafeReport(numbers)
		if isSafe {
			safeCount++
		}
		
		// Check if the report is safe with the second set of requirements
		if isSafe || isKindaSafeReport(numbers) {
			safeCount2++
		}
	}

	// Check for errors while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return safeCount, safeCount2
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

func isKindaSafeReport(numbers []int) bool {
	// Try removing each level one at a time and check if the result is safe
	for i := 0; i < len(numbers); i++ {
		newNumbers := append([]int{}, numbers[:i]...)  // Append the numbers before index i to an empty slice
		newNumbers = append(newNumbers, numbers[i+1:]...)  // Append the part after index i to the slice above
		if isSafeReport(newNumbers) {
			return true
		}
	}

	return false
}

func main() {
	startTime := time.Now()

	safeReports, safeReports2 := FindSafeReportsFromFile()
	fmt.Printf("Total safe reports: %d, including the kinda safe: %d\n", safeReports, safeReports2)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
