package util

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

// FindFreeTCPPort finds a free TCP port in the dynamic port allocation range
func FindFreeTCPPort() int {

	var port int
	for {
		// Dynamic port allocations should be in the range 49152-65535
		port = 49152 + RandomIntN(16383)

		listener, err := net.Listen(`tcp`, ":"+strconv.Itoa(port))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't listen on port %d: %s", port, err)
			continue
		}
		listener.Close()
		break
	}
	return port
}
