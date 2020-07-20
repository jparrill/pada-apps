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
	status, _ := CheckEth(ifName)
	fmt.Printf("%s", status)
}

func getInterface(iface_status string) string {
	ifaces, _ := net.Interfaces()
	var result string

	for _, i := range ifaces {
		addrs, _ := i.Addrs()

		if iface_status == "disconnected" {
			if len(addrs) < 1 {
				//fmt.Printf("DISCONNECTED %v: %v\n", i.Name, addrs)
				result = i.Name
			}

		} else if iface_status == "connected" {
			if len(addrs) > 0 {
				//fmt.Printf("CONNECTED %v: %s\n", i.Name, addrs[0])
				result = i.Name
			}

		} else {
			panic("Case not covered")
		}
	}

	return result
}

func CheckEth(ifName string) (string, int) {
	byNameInterface, err := net.InterfaceByName(ifName)
	var rc int

	if err != nil {
		rc = 2
		return utils.Warn(eth_icon), rc
	}

	addresses, _ := byNameInterface.Addrs()

	if len(addresses) > 0 {
		rc = 0
		return utils.Info(eth_icon), rc
	} else {
		rc = 1
		return utils.Erro(eth_icon), rc
	}
}
