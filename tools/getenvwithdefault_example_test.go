package tools

import (
	"flag"
	"fmt"
)

func ExampleGetenvWithDefault() {
	var boolOption bool
	var intOption int

	flag.IntVar(&intOption, "intoption", GetenvWithDefault("MY_INT_OPTION", 10), "an example int option")
	flag.BoolVar(&boolOption, "booloption", GetenvWithDefault("MY_BOOL_OPTION", true), "an example bool option")
	flag.Parse()

	fmt.Printf("int option: %d\n", intOption)
	fmt.Printf("bool option: %t\n", boolOption)
}
