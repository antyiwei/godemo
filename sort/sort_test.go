package main

import (
	"fmt"
	"testing"

	"github.com/psilva261/timsort"
)

type Record struct {
	ssn  int
	name string
}

func BySsn(a, b interface{}) bool {
	return a.(Record).ssn < b.(Record).ssn
}

func ByName(a, b interface{}) bool {
	return a.(Record).name < b.(Record).name
}

func TestTimesort(t *testing.T) {
	db := make([]interface{}, 3)
	db[0] = Record{123456789, "joe"}
	db[1] = Record{101765430, "sue"}
	db[2] = Record{345623452, "mary"}

	// sorts array by ssn (ascending)
	timsort.Sort(db, BySsn)
	fmt.Printf("sorted by ssn: %v\n", db)

	// now re-sort same array by name (ascending)
	timsort.Sort(db, ByName)
	fmt.Printf("sorted by name: %v\n", db)
}
