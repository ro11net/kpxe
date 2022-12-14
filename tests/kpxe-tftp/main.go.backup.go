package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

const (
	// The default port number for the TFTP server.
	defaultPort = "69"
)

func main() {
	// Bind to the default TFTP port.
	l, err := net.Listen("udp", "127.0.0.1:"+defaultPort)
	if err != nil {
		fmt.Println("Error binding to port:", err)
		return
	}
	defer l.Close()

	fmt.Println("TFTP server listening on port", defaultPort)

	// Loop forever, accepting and handling TFTP requests.
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the request in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// TODO: Implement handling of TFTP requests here.
	fmt.Println("Received TFTP request from", conn.RemoteAddr())

	// Close the connection when the function returns.
	defer conn.Close()
}




########

package main

import (
"errors"
"fmt"
"io"
"net/http"
"os"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
