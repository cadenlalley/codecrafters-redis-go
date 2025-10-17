package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	testCases := []struct {
		name             string
		command          Command
		expectedResponse string
		shouldError      bool
	}{
		{
			name: "valid echo",
			command: Command{
				Command:   "ECHO",
				Arguments: []string{"hey"},
			},
			expectedResponse: "$3\r\nhey\r\n",
			shouldError:      false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			response := testCase.command.Run()

			assert.Equal(t, testCase.expectedResponse, response)
		})
	}
}
