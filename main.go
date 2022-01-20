/*
 * This program provides a set of commands for manipulating mailman3 lists
 */

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagc string // location of the configuration file
	flagb bool   // set true to use JSON booleans, set false to use JSON strings
)

func init() {
	flag.StringVar(&flagc, "c", "./mailman3test.cfg", "configuration file name")
	flag.BoolVar(&flagb, "b", false, "use JSON booleans, not strings")
	flag.Parse()
}

func usage() {
	fmt.Println("This program uses the REST API to subscribe a user to a list")
	fmt.Println("Usage:")
	fmt.Println("mailman3test <flags> <list> <address>")
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	config()
	subscribeCmd()
}
