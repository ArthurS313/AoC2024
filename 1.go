package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Create2SortedSlices() ([]int, []int) {
	f, err := os.Open("1.txt")
    if err != nil {
        fmt.Println(err)}
	r := bufio.NewReader(f)
	a := []int{}
	b := []int{}
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


func main() {
	a,b:=Create2SortedSlices()

	sum, err := AbsDifference(a, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(sum)
}