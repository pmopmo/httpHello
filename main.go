// A small program tha listens on a ports
// and write a very short html response
// I used this to test/verify a load balancing setup

// Takes one integer (ports nr) between 1025 & 65535 as an argument

package main

import (
	"fmt"
	"os"

	"httpHello/sayhello"
)

func main() {
	usage := "Expects one or more integers between 1025 & 65535 as argument."

	// Any arguments?
	if len(os.Args) < 2 {
		fmt.Printf("%s", usage)
		_, _ = fmt.Fprint(os.Stderr, usage)
		os.Exit(1)

	} else {

		// TODO: if more than one argument loop over them
		err := sayhello.StartListener(os.Args[1])
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, err)
		}
	}
} // end of main
