package twenty_test

import (
	"testing"
	twenty "v2/20"
	"v2/common"
)

func Test_Reverse(t *testing.T) {
	tests := []struct {
		str      string
		expected string
	}{
		{
			"snow dog sun",
			"sun dog snow",
		},
		{
			"nice dog you bastard",
			"bastard you dog nice",
		},
		{
			"",
			"",
		},
		{
			"foo",
			"foo",
		},
	}

	for _, test := range tests {
		common.Equal(t, twenty.Reverse(test.str), test.expected)
	}
}
