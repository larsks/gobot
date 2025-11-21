package tools

import (
	"log"
)

func ExampleRangeIterator() {
	for num, err := range RangeIterator("1,4,6-10") {
		if err != nil {
			log.Fatalf("bad range specification")
		}
		println(num)
	}

	// Prints:
	// 1
	// 4
	// 6
	// 7
	// 8
	// 9
	// 10
}
