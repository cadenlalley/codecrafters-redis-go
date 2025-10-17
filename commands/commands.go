package commands

import (
	"errors"
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/resp"
)

type Command struct {
	Command   string
	Arguments []string
}

const (
	PING = "PING"
	ECHO = "ECHO"
)

func (c Command) Run() (string, error) {
	switch c.Command {
	case PING:
		return ping()
	case ECHO:
		return echo(c.Arguments...)
	}

	return "command not found", nil
}

func ping() (string, error) {
	return fmt.Sprintf(resp.SIMPLESTRING, "PONG"), nil
}

func echo(args ...string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("echo must have arguements")
	}

	return fmt.Sprintf(resp.BULKSTRING, len(args[0]), args[0]), nil
}
