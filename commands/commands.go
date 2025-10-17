package commands

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/cache"
	"github.com/codecrafters-io/redis-starter-go/resp"
)

type Command struct {
	Command   string
	Arguments []string
}

const (
	PING = "PING"
	ECHO = "ECHO"
	SET  = "SET"
	GET  = "GET"
)

func (c Command) Run() string {
	switch c.Command {
	case PING:
		return ping()
	case ECHO:
		return echo(c.Arguments...)
	case SET:
		return set(c.Arguments...)
	case GET:
		return get(c.Arguments...)
	}

	return "command not found"
}

func ping() string {
	return fmt.Sprintf(resp.SIMPLESTRING, "PONG")
}

func echo(args ...string) string {
	if len(args) == 0 {
		return ""
	}

	return fmt.Sprintf(resp.BULKSTRING, len(args[0]), args[0])
}

func set(args ...string) string {
	err := cache.Set(args[0], args[1])
	if err != nil {
		return "IDK"
	}

	return resp.GetOKResponse()
}

func get(args ...string) string {
	val, err := cache.Get(args[0])
	if err != nil {
		return "IDK"
	}

	return fmt.Sprintf(resp.SIMPLESTRING, val)
}
