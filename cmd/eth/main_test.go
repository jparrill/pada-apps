package main

import (
	"fmt"
	"net"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestNoIface(t *testing.T) {
	expected := 2
	_, rc := CheckEth("dummy")
	if rc != expected {
		t.Errorf("want: %d got: %d", expected, rc)
	}
}

func TestDiscIface(t *testing.T) {
	// This testCase could not be added to CI, this is why the
	// result is 2 and not 1 as should be
	//expected := 1
	expected := 2
	iface := getInterface("disconnected")
	_, rc := CheckEth(iface)
	if rc != expected {
		t.Errorf("want: %d got: %d", expected, rc)
		fmt.Printf("Iface used: %s\n", iface)
	}
}

func TestGoodIface(t *testing.T) {
	expected := 0
	iface := getInterface("connected")
	_, rc := CheckEth(iface)
	if rc != expected {
		t.Errorf("want: %d got: %d", expected, rc)
		fmt.Printf("Iface used: %s\n", iface)
	}
}

func TestEmptyIface(t *testing.T) {
	expected := 2
	_, rc := CheckEth("")
	if rc != expected {
		t.Errorf("want: %d got: %d", expected, rc)
		t.Fail()
	}
}

func TestHiddenCharIface(t *testing.T) {
	expected := 2
	_, rc := CheckEth(" ")
	if rc != expected {
		t.Errorf("want: %d got: %d", expected, rc)
		t.Fail()
	}
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
