package human_logfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFailToRowFileIfValuesAreMissing(t *testing.T) {

	var tests = []struct {
		whoKilled string
		whoDied   string
		meanOfDth string
		want      error
	}{
		{"", "", "", ErrMissingFields},
		{"Gom", "Xyz", "", ErrMissingFields},
		{"", "Frank", "Joel", ErrMissingFields},
		{"", "Frank", "", ErrMissingFields},
		{"", "", "John", ErrMissingFields},
		{"Mark", "", "Frank", ErrMissingFields},
	}

	for _, test := range tests {
		err := NewHumanLogFileRow(test.whoKilled, test.whoDied, test.meanOfDth).Validation()
		assert.Equal(t, err, test.want, "Expected: %d - Got: %d", test.want, err)
	}

}
