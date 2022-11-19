package timeparse

import "testing"

func TestParseTime(t *testing.T) {

	table := []struct {
		time string
		ok   bool
	}{
		{"19:02:02", true},
		{"29:02:02", false},
		{"hello", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		// t.Logf("Testing for %v", data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v error should be nil", data.time, err)
		}
	}

}
