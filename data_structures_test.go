package zerosixty_test

import (
	"strings"
	"testing"
)

func TestArrayTypeWithRangeOperator(t *testing.T) {
	// arrays fixed size
	anArray := [4]string{"Asynchrony", "World", "Wide", "Technology"}

	t.Log("anArray content:")
	for idx, value := range anArray {
		t.Logf("\tid: %d, value: %v\n", idx, value)
	}

	// memory allocation done at assignment
	// zero values assigned for any index unassigned by coder
	unFilledArray := [4]string{"Missed", "By", "1"}

	t.Log("unFilledArray content:")
	for idx, value := range unFilledArray {
		t.Logf("\tid: %d, value: %v\n", idx, value)
	}

	unFilledArray[3] = "covfefe"
	t.Logf("\tid: %d, value: %v\n", 4, unFilledArray[3])
}

func TestSliceTypeWithRangeOperator(t *testing.T) {
	/*
	   slice is 3 word data structure
	   1st word: points to backing array
	   2nd word: length of slice
	   3rd word: capacity of slice

	   slice is basically a dynamic array, able to change capacity at runtime
	   initializes like array without indicating size
	   slice is a reference type

	   slice is Go's most important data structure
	*/

	strSlice := []string{"like", "arrays", "only", "better"}

	t.Log("strSlice:")
	for idx, value := range strSlice {
		t.Logf("\tid: %d, value: %v\n", idx, value)
	}

	// can extend with append
	strSlice = append(strSlice, "and", "extendable")

	t.Logf("strSlice after append:")
	for idx, value := range strSlice {
		t.Logf("\tid: %d, value: %v\n", idx, value)
	}
}

func TestSliceReassignments(t *testing.T) {
	slice := make([]string, 0, 0) // initializes slice with 0 length and 0 capacity

	slice = append(slice, "with", "great", "power", "comes", "great", "responsibility")
	t.Log(strings.Join(slice, " "))

	newSlice := slice[:] // returns copy of original slice
	newSlice[2] = "food"
	t.Log(strings.Join(newSlice, " "))

	newerSlice := append(slice[:len(newSlice)-1], "indigestion")
	t.Log(strings.Join(newerSlice, " "))
}

func TestMapTypeRangeOperator(t *testing.T) {
	/*
			map is a k/v data structure
			hash function done for you (super secure, google designed)
		    map is a reference type
	*/
	aMap := map[string][]string{
		"marvel": {"spiderman", "hulk", "ironman"},
		"dc":     {"batman", "superman", "wonderwoman"},
	}

	t.Log("amap:")
	for k, v := range aMap {
		t.Logf("\t%s", k)
		for _, hero := range v {
			t.Logf("\t\t%s", hero)
		}
	}

	delete(aMap, "dc") // drops k/v pair

	t.Log("aMap dc removed:")
	for k, v := range aMap {
		t.Logf("\t%s", k)
		for _, hero := range v {
			t.Logf("\t\t%s", hero)
		}
	}

	aMap["marvel"] = append(aMap["marvel"], "thor", "black widow")

	t.Log("aMap addition:")
	for k, v := range aMap {
		t.Logf("\t%s", k)
		for _, hero := range v {
			t.Logf("\t\t%s", hero)
		}
	}
}

func TestZeroTypesOfReferenceTypes(t *testing.T) {
	// reference types all have a zero value of nil
	var zeroSlice []bool
	var zeroMap map[string]bool

	if zeroSlice == nil {
		t.Logf("zeroSlice == nil")
	}

	if zeroMap == nil {
		t.Logf("zeroMap == nil")
	}
}
