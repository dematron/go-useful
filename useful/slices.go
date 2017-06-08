package useful

import (
	"fmt"
	"reflect"
	"sort"
)

func removeDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

func sortStringSlice(result []string) []string {
	sort.Strings(result)
	return result
}

// http://www.tapirgames.com/blog/golang-slice-comparison
func compareIntSlices_BCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)] // this line is the key
	for i, v := range a {
		if v != b[i] { // here is no bounds checking for b[i]
			return false
		}
	}

	return true
}

func compareSlices_BCE2(a, b interface{}, lena, lenb int) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	aa := make([]interface{}, lena)
	bb := make([]interface{}, lenb)

	switch a.(type) {
	case []int:
		for i, d := range a.([]int) {
			aa[i] = d
		}
		for i, d := range b.([]int) {
			bb[i] = d
		}
	case []float64:
		for i, d := range a.([]float64) {
			aa[i] = d
		}
		for i, d := range b.([]float64) {
			bb[i] = d
		}
	case []string:
		for i, d := range a.([]string) {
			aa[i] = d
		}
		for i, d := range b.([]string) {
			bb[i] = d
		}
	default:
		// And here I'm feeling dumb. ;)
		fmt.Printf("I don't know what to compare.")
		return false
	}

	if len(aa) != len(bb) {
		return false
	}

	if (aa == nil) != (bb == nil) {
		return false
	}

	bb = bb[:len(aa)] // this line is the key
	for i, v := range aa {
		if v != bb[i] { // here is no bounds checking for b[i]
			return false
		}
	}

	return true
}

func SliceExamples() {
	fmt.Println("Examples of work with slices.")
	fmt.Println()
	elementsint := []int{100, 200, 300, 100, 200, 400, 0}
	fmt.Println("Int slice with duplicates:", elementsint)

	// Test our method.
	resultint := removeDuplicates(elementsint)
	fmt.Println("Int slice without duplicates: ", resultint)

	elementsstring := []string{"cat", "dog", "cat", "bird"}
	fmt.Println("String slice with duplicates:", elementsstring)

	// Remove string duplicates, ignoring order.
	resultstring := removeDuplicatesUnordered(elementsstring)
	fmt.Println("String slice without duplicates:", resultstring)

	//Sorting string slice
	resultsortstring := sortStringSlice(resultstring)
	fmt.Println("Sorted string slice without duplicates:", resultsortstring)

	// Slice compare
	elementsstring2 := []string{"dog", "cat", "cat", "bird"}
	fmt.Println("Comparing slice", elementsstring, "and", elementsstring)
	if compareSlices_BCE2(elementsstring, elementsstring, len(elementsstring), len(elementsstring)) {
		fmt.Println("First and Second slice are equal")
	} else {
		fmt.Println("First and Second slice are NOT equal")
	}

	fmt.Println("Comparing slice", elementsstring, "and", elementsstring2)
	if compareSlices_BCE2(elementsstring, elementsstring2, len(elementsstring), len(elementsstring2)) {
		fmt.Println("First and Second slice are equal")
	} else {
		fmt.Println("First and Second slice are NOT equal")
	}
}
