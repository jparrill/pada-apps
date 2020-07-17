package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetCoreTemp(t *testing.T) {
	expected := 0
	_, rc := getCoreTemp()
	if rc != expected {
		t.Errorf("want: %d got: %d", expected, rc)
	}
}

func TestIslmSensorsInstalled(t *testing.T) {
	expected := 0
	_, rc := islmSensorsInstalled()
	if rc != expected {
		t.Errorf("want: %d got: %d", expected, rc)
	}
}
