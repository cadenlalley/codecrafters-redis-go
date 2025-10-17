package parser

import (
	"strings"

	"github.com/codecrafters-io/redis-starter-go/commands"
)

const (
	ARRAY = "*"
)

func Decode(input []byte) commands.Command {
	inputString := string(input)
	parts := strings.Split(inputString, "\r\n")

	var trimmedParts []string
	for i := 2; i < len(parts); i += 2 {
		trimmedParts = append(trimmedParts, parts[i])
	}

	return commands.Command{
		Command:   trimmedParts[0],
		Arguments: trimmedParts[1:],
	}
}
