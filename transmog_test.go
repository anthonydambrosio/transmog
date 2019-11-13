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

func TestGetBool(t *testing.T) {
	v, _ := tm.Get([]string{"k8"})
	b, err := strconv.ParseBool(v)
	if !b {
		t.Errorf("Key k8 is not true: %v", err)
	}
}

func TestSetBool(t *testing.T) {
	err := tm.Set([]string{"k8"}, "false")
	if err != nil {
		t.Error(err)
	}
}

func TestAddStringProperty(t *testing.T) {
	tm.Set([]string{"k9"}, "Lassie")
	v, _ := tm.Get([]string{"k9"})
	if v != "Lassie" {
		t.Error("k9 is not Lassie")
	}
}

func TestAddFloatProperty(t *testing.T) {
	tm.Set([]string{"k9"}, "3.14")
	v, _ := tm.Get([]string{"k9"})
	if v != "3.14" {
		t.Error("Value is not 3.14")
	}
	if f, _ := strconv.ParseFloat(v, 64); f != 3.14 {
		t.Error("Float is not float")
	}
}

func TestGetXml(t *testing.T) {
	var tmog Transmog
	err := tmog.LoadFile("test.xml")
	if err != nil {
		t.Error(err)
	}
	value, _ := tmog.Get([]string{"address", "contact", "name"})
	if value != "Tanmay Patil" {
		t.Errorf("name in text.xml is not 'Tanmay Patil', is %v", value)
	}
}

func TestSetXml(t *testing.T) {
	var tmog Transmog
	err := tmog.LoadFile("test.xml")
	if err != nil {
		t.Error(err)
	}
	err = tmog.Set([]string{"address", "contact", "name"}, "John Smith")
	if err != nil {
		t.Error(err)
	}
}

func TestMain(m *testing.M) {
	err := tm.LoadFile("test.json")
	if err != nil {
		fmt.Print(fmt.Errorf("%v", err.Error()))
		os.Exit(1)
	}
	os.Exit(m.Run())
}
