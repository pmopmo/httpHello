// A small program tha listens on a ports
// and write a very short html response
// I used this to test/verify a load balancing setup

// Takes one integer (ports nr) between 1025 & 65535 as an argument

package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/pmopmo/httpHello/sayhello"
)

func main() {
	usage := "Expects one or more integers between 1025 & 65535 as argument."

	// Any arguments?
	if len(os.Args) < 2 {
		fmt.Printf("%s", usage)
		_, _ = fmt.Fprint(os.Stderr, usage)
		os.Exit(1)

	} else {
		var wg sync.WaitGroup

		// Had to to this here got a panic for multiple handlers for "/" when I had it in StartListener
		// TODO find a better way if any
		sayhello.SetHandler()

		// os.Args[1:] => everything except first element
		for _, port := range os.Args[1:] {
			wg.Add(1)
			go func(p string) { _ = sayhello.StartListener(p) }(port)
		}
		wg.Wait()
	}
} // end of main
