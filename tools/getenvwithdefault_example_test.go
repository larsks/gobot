package tools

import (
	"fmt"
	"time"
)

func ExampleGetenvWithDefault() {
	var boolOption bool
	var intOption int
	var timeOption time.Time

	defaultTime, _ := time.Parse("2006-01-01", "2000-10-10")

	intOption = GetenvWithDefault("MY_INT_OPTION", 10)
	boolOption = GetenvWithDefault("MY_BOOL_OPTION", true)
	timeOption = GetenvWithDefault("MY_TIME_OPTION", defaultTime, "2006-01-02")

	fmt.Printf("int var is %d, bool var is %t, time var is %s", intOption, boolOption, timeOption)
}
