package util

import (
	"testing"
)

func TestRandString(t *testing.T) {
	for i := 10; i <= 20; i++ {
		s1 := RandString(i)
		s2 := RandString(i)
		if len(s1) != len(s2) || len(s1) != i {
			t.Errorf("RandString created string of wrong length: (s1,s2,i)=(%i,%i,%i)\n", len(s1), len(s2), i)
		}
		if s1 == s2 {
			t.Errorf("RandString created two identical strings:\n\t%v\n\t%v\n", s1, s2)
		}
	}
	if RandString(0) != "" {
		t.Errorf("RandString(0) != \"\"\n")
	}
	if RandString(-1) != "" {
		t.Errorf("RandString(-1) != \"\"\n")
	}
}
