package ports

import (
	"fmt"
	"net"
	"os"
)

func Available(port string) bool {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return false
	}

	err = ln.Close()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Couldn't stop listening on ports %q: %s", port, err)
		return false //os.Exit(1)
	}

	return true
}
