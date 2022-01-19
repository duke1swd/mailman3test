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
)

func init() {
	flag.StringVar(&flagc, "c", "/opt/mailman/mm/mm3util.cfg", "configuration file name")
	flag.Parse()
}

func usage() {
	fmt.Println("Flags:")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	config()
	subscribeCmd()
}
