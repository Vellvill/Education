package twelve_test

import (
	"testing"
	twelve "v2/12"
	"v2/common"
)

func Test_Solution(t *testing.T) {
	tests := []struct {
		str      []string
		expected []string
	}{
		{
			[]string{"cat", "cat", "dog", "cat", "tree"},
			[]string{"cat", "dog", "tree"},
		},
		{
			[]string{"cat", "dog"},
			[]string{"cat", "dog"},
		},
		{
			[]string{"cat", "dog", "dog", "cat"},
			[]string{"cat", "dog"},
		},
		{
			[]string{"cat"},
			[]string{"cat"},
		},
		{
			[]string{"dog", "cat", "elephant", "junior", "chicken"},
			[]string{"dog", "cat", "elephant", "junior", "chicken"},
		},
	}

	for _, test := range tests {
		common.Equal(t, twelve.Solution(test.str), test.expected)
	}
}
