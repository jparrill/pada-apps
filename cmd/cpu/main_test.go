package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetCPUs(t *testing.T) {
	expected := 0
	_, rc := getCPUs()
	if rc != expected {
		t.Errorf("want: %d got: %d", expected, rc)
	}
}

func TestLoadAvgWrongFieldGt5(t *testing.T) {
	expected := 2
	var i int64 = 6
	for ; i <= 10; i++ {
		_, rc := getLoadAvg(i)
		if rc != expected {
			t.Errorf("want: %d got: %d", expected, rc)
		}
	}
}

func TestLoadAvgWrongFieldGt3(t *testing.T) {
	expected := 1
	var i int64 = 4
	for ; i <= 5; i++ {
		_, rc := getLoadAvg(i)
		if rc != expected {
			t.Errorf("want: %d got: %d", expected, rc)
		}
	}
}

func TestLoadAvgCorrectFields(t *testing.T) {
	expected := 0
	var i int64 = 1
	for ; i <= 3; i++ {
		_, rc := getLoadAvg(i)
		if rc != expected {
			t.Errorf("want: %d got: %d", expected, rc)
		}
	}
}
