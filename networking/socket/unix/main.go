package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
)

const (
	network = "unix"
	host    = "localhost"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("missing required arguments: <socket> <uri>")
		os.Exit(1)
	}

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial(network, os.Args[1])
			},
		},
	}

	resp, err := client.Get(fmt.Sprintf("http://%s", path.Join(host, os.Args[2])))
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Printf("error: %v", err)
	}
}
