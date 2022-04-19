package tventy_six_test

import (
	"testing"
	tventy_six "v2/26"
	"v2/common"
)

func Test_isUnique(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{
			"abcd",
			true,
		},
		{
			"abCdefAaf",
			false,
		},
		{
			"aabcd",
			false,
		},
		{
			"qwertyuiopQWERTYUIOP",
			false,
		},
		{
			"qwertyuiop",
			true,
		},
	}

	for _, test := range tests {
		common.Equal(t, tventy_six.IsUnique(test.str), test.expected)
	}
}
