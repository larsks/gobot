package tools

import (
	"time"

	flag "github.com/spf13/pflag"
)

func ExampleGetenvWithDefault() {
	var boolOption bool
	var intOption int
	var timeOption time.Time

	defaultTime, _ := time.Parse("2006-01-02", "1900-01-01")

	// Note that we are using pflag (https://pkg.go.dev/github.com/spf13/pflag) in this example
	// because it supports more flag types than the stock flag module.
	flag.IntVar(&intOption, "intoption", GetenvWithDefault("MY_INT_OPTION", 10), "an example int option")
	flag.BoolVar(&boolOption, "booloption", GetenvWithDefault("MY_BOOL_OPTION", true), "an example bool option")
	flag.TimeVar(&timeOption, "timeoption",
		// We need to provide a time format for parsing the environment variable
		// as a third parameter to ExampleGetenvWithDefault.
		GetenvWithDefault("MY_TIME_OPTION", defaultTime, "2006-01-02"),

		// pflag also needs a list of time formats for parsing command line options
		[]string{"2006-01-02"},
		"an example time option")
	flag.Parse()
}
