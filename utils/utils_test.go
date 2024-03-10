package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	cases := []struct {
		testName string
	}{
		{
			testName: "1. Positive test",
		},
	}

	for _, c := range cases {
		t.Logf(CurrentlyTesting, c.testName)
		password := GenerateRandomString(7)
		assert.NotNil(t, password)
	}
}
