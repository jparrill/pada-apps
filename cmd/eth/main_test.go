package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/jparrill/pada-apps/pkg/utils"
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
	utils.SkipCI(t)
	expected := 1
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
