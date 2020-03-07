package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range interfaces {

		up := "UP"
		if i.Flags & net.FlagUp == 0 {
			up = "DOWN"
		}
		fmt.Printf("%2d %-4s %-8s %18s MTU %5d %s\n", i.Index, up, i.Name, i.HardwareAddr, i.MTU, i.Flags)

		addrs, err := i.Addrs()
		if err != nil {
			fmt.Printf("   unicast interface addresses: %v\n", err)
			continue
		}

		if len(addrs) > 0 {
			fmt.Println("   unicast interface addresses:")
			for _, a := range addrs {
				fmt.Printf("     %s\n", a)
			}
		}
	}
}
