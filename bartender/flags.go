package main

import "strings"

/* Defining an array of Flags to allow multiple values:
Example:
	go run main.go --list1 value1 --list1 value2
*/
type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ",")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}
