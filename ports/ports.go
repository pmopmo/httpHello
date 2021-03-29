package ports

import (
	"fmt"
	"net"
	"os"
)

// Checks if a port is free by trying to listen to it and then release it
// I'd really like to find a better way to do this
func Available(port string) bool {

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return false
	}

	err = ln.Close()

	// post completion notnil
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Couldn't stop listening on ports %q: %s", port, err)
		return false //os.Exit(1)
	}

	return true
}
