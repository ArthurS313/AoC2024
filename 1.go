package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Create2SortedSlices() ([]int, []int) {
	// Opens file containing the two columns
	f, err := os.Open("1.txt")
    if err != nil {
        fmt.Println(err)}
	r := bufio.NewReader(f)
	// a is a list that contains all the ints on the left
	a := []int{}
	// b is a list that contains all the ints on the right
	b := []int{}
	// For each line, remove the linebreak character, splits the line by the triple space marker, and adds the two numbers to their respective slice
	for {
	    line, err := r.ReadString('\n')
	    if err != nil {
			break
	    }
		var aString, bString, _ = strings.Cut(line, "   ")
		aInt, _  := strconv.Atoi(aString)
		a= append(a, aInt)
		bString=bString[:len(bString)-2]
		bInt, _  := strconv.Atoi(bString)
		b= append(b, bInt)
		
	}
	defer f.Close()
	// Sort the two slices
	sort.Ints(a)
	sort.Ints(b)
	return a,b
}

func AbsDifference(a, b []int) (int, error) {
	// Ensure the slices are of the same size.
	if len(a) != len(b) {
		return 0, fmt.Errorf("slices must be of the same size")
	}

	// Create a slice to store the results.
	var sum int

	// Calculate the absolute difference for each pair of elements.
	for i := 0; i < len(a); i++ {
		sum+= int(math.Abs(float64(a[i] - b[i])))
	}

	return sum, nil
}

func CalculateSimilarity(a []int, b []int) int {
	// Create a map to count occurrences of numbers in the right list
	bCounts := make(map[int]int)
	for _, num := range b {
		bCounts[num]++
	}

	// Calculate the similarity score
	similarityScore := 0
	for _, num := range a {
		similarityScore += num * bCounts[num]
	}

	return similarityScore
}


func main1() {
	startTime := time.Now()

	a,b:=Create2SortedSlices()
	//sum1 is the result for the first part of day 1
	sum1, err := AbsDifference(a, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//sum2 is the result for the second part of day 2
	sum2:= CalculateSimilarity(a,b)
	fmt.Println(sum1, sum2)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}