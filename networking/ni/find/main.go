package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
)

var ip = "127.0.0.1"

func main() {

	args := os.Args[1:]
	if len(args) > 0 {
		ip = args[0]
	}

	ni, ipNet, err := findNetworkInterface(net.ParseIP(ip))
	if err != nil {
		log.Fatalf("cannot find network inteface for %s ip", ip)
	}
	fmt.Printf("ip: %s\n", ip)
	fmt.Printf("newtwork interface: %-8s %18s MTU %5d %s\n", ni.Name, ni.HardwareAddr, ni.MTU, ni.Flags)
	fmt.Printf("address: %s ip: %s mask: %s\n", ipNet, ipNet.IP, ipNet.Mask)
}

func findNetworkInterface(dstIP net.IP) (net.Interface, *net.IPNet, error) {

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range interfaces {

		if i.Flags&net.FlagUp == 0 {
			continue
		}

		addrs, err := i.Addrs()
		if err != nil {
			return net.Interface{}, nil, err
		}

		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok {
				if ipnet.Contains(dstIP) {
					return i, ipnet, nil
				}
			}
		}
	}
	return net.Interface{}, nil, errors.New("no interface found")
}
