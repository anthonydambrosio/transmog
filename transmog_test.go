package transmog

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

var tm Transmog

func TestGetk1(t *testing.T) {
	v, _ := tm.Get([]string{"k1"})
	if v != "Key One" {
		t.Errorf("Key k1 is not 'Key One'")
	}
}

func TestGetk2(t *testing.T) {
	v, _ := tm.Get([]string{"k2"})
	if v != "2" {
		t.Errorf("Key k2 is not '2'")
	}
}

func TestSetk1(t *testing.T) {
	_ = tm.Set([]string{"k1"}, "Kay 1")
	v, _ := tm.Get([]string{"k1"})
	if v != "Kay 1" {
		t.Errorf("Key k2 is not 'Kay 1'")
	}
}

func TestGetk8(t *testing.T) {
	v, _ := tm.Get([]string{"k8"})
	b, err := strconv.ParseBool(v)
	if !b {
		t.Errorf("Key k8 is not true: %v", err)
	}
}

func TestMain(m *testing.M) {
	err := tm.Load("test.json")
	if err != nil {
		fmt.Print(fmt.Errorf("%v", err.Error()))
		os.Exit(1)
	}
	os.Exit(m.Run())
}
