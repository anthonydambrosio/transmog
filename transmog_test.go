package transmog

import (
	"testing"
)

func TestLoad(t *testing.T) {
	var tm Transmog
	err := tm.Load("test.json")
	if err != nil {
		t.Errorf("Failed to Load test.json: %v", err)
	}
}

func TestGetk1(t *testing.T) {
	var tm Transmog
	_ = tm.Load("test.json")
	v, _ := tm.Get([]string{"k1"})
	if v != "Key One" {
		t.Errorf("Key k1 is not 'Key One'")
	}

}
