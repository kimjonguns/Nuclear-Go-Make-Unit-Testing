package tabletest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListNama(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "funcHafi",
			Request:  "Hafi",
			Expected: "Hafi",
		},
		{
			Name:     "funcIhza",
			Request:  "Ihza",
			Expected: "Ihza",
		},
		{
			Name:     "funcFarhana",
			Request:  "Farhana",
			Expected: "Farhana",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := ListNama(test.Request)
			assert.Equal(t, test.Expected, result)
		})
	}
}
