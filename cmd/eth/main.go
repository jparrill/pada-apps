package main

import (
	"fmt"
	"net"
	"os"

	"github.com/jparrill/pada-apps/pkg/utils"
)

var (
	eth_icon = ""
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <interface name>\n", os.Args[0])
		fmt.Printf("Interfaces: ")
		interfaces, _ := net.Interfaces()

		for _, i := range interfaces {
			fmt.Printf("%v ", i.Name)
		}
		os.Exit(0)
	}

	ifName := os.Args[1]

	byNameInterface, err := net.InterfaceByName(ifName)

	if err != nil {
		fmt.Printf("No Int: %s %s", ifName, utils.Yellow(eth_icon))
		os.Exit(0)
	}

	addresses, _ := byNameInterface.Addrs()

	if len(addresses) > 0 {
		fmt.Printf("%s", utils.Info(eth_icon))
	} else {
		fmt.Printf("%s", utils.Erro(eth_icon))
	}
}
