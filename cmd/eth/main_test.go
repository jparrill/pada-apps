package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestNoIface(t *testing.T) {
	expected := 2
	_, rc := CheckEth("ens5")
	if rc != expected {
		t.Errorf("want: %d got: %d", expected, rc)
		t.Fail()
	}
}

func TestDiscIface(t *testing.T) {
	expected := 1
	_, rc := CheckEth("ens4")
	if rc != expected {
		t.Errorf("want: %d got: %d", expected, rc)
		t.Fail()
	}
}

func TestGoodIface(t *testing.T) {
	expected := 0
	_, rc := CheckEth("wlp2s0")
	if rc != expected {
		t.Errorf("want: %d got: %d", expected, rc)
		t.Fail()
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
