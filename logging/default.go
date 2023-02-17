package logging

import "log"

// DefaultDebugPrint defines a default function that prints resultant query and arguments before being executed
// and when the Debug flag is true
func DefaultDebugPrint(q string, args []interface{}, err error) {
	if q != "" {
		log.Printf("query: %v \nargs: %v\n", q, args)
	}

	if err != nil {
		log.Println("err: ", err.Error())
	}
}
