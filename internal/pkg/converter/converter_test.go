package converter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToBool(t *testing.T) {
	testCases := []struct {
		value    string
		expected bool
	}{
		{
			value:    "true",
			expected: true,
		},
		{
			value:    "1",
			expected: true,
		},
		{
			value:    "tRuE",
			expected: true,
		},
	}
	for _, testCase := range testCases {
		desc := fmt.Sprintf("Expected %s to be %v", testCase.value, testCase.expected)
		assert.Equal(t, testCase.expected, StringToBool(testCase.value), desc)
	}
}
