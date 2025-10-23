package commands

import (
	"strconv"
	"strings"
	"time"

	"github.com/codecrafters-io/redis-starter-go/cache"
	"github.com/codecrafters-io/redis-starter-go/resp"
)

type Command struct {
	Command   string
	Arguments []string
}

const (
	PING  = "ping"
	ECHO  = "echo"
	SET   = "set"
	GET   = "get"
	RPUSH = "rpush"
)

func (c Command) Run() string {
	switch strings.ToLower(c.Command) {
	case PING:
		return ping()
	case ECHO:
		return echo(c.Arguments...)
	case SET:
		return set(c.Arguments...)
	case GET:
		return get(c.Arguments...)
	case RPUSH:
		return listPush(c.Arguments...)
	default:
		return resp.NewErrorResponse("command not found")
	}
}

func ping() string {
	return resp.NewSimpleString("PONG")
}

func echo(args ...string) string {
	if len(args) == 0 {
		return resp.NewErrorResponse("wrong number of arguments for the echo command")
	}

	return resp.NewBulkString(args[0])
}

func set(args ...string) string {
	if len(args) < 2 || len(args)%2 != 0 {
		return resp.NewErrorResponse("wrong number of arguments for the set command")
	}

	cacheItem := cache.Item{
		Value: args[1],
	}

	if len(args) > 2 {
		exp, err := strconv.Atoi(args[3])
		if err != nil {
			return resp.NewErrorResponse("expiration time is not an integer value")
		}
		switch strings.ToLower(args[2]) {
		case "ex":
			time.AfterFunc((time.Second * time.Duration(exp)), func() {
				cache.Delete(args[0])
			})
		case "px":
			time.AfterFunc((time.Millisecond * time.Duration(exp)), func() {
				cache.Delete(args[0])
			})
		}
	}

	cache.Set(args[0], cacheItem)

	return resp.NewOKResponse()
}

func get(args ...string) string {
	if len(args) == 0 {
		return resp.NewErrorResponse("wrong number of arguments for the get command")
	}

	val, ok := cache.Get(args[0])
	if !ok {
		return resp.GetNullBulkString()
	}

	return resp.NewSimpleString(val.Value)
}

func listPush(args ...string) string {
	if len(args) < 2 {
		return resp.NewErrorResponse("wrong number of arguments for the rpush command")
	}

	item, ok := cache.Get(args[0])
	if !ok {
		newItem := cache.Item{
			Value: []any{args[1:]},
		}
		cache.Set(args[0], newItem)

		return resp.NewUnsignedInteger(1)
	}

	slice, ok := item.Value.([]any)
	if !ok {
		return resp.NewErrorResponse("value is not a list")
	}
	slice = append(slice, args[1])

	item.Value = slice

	return resp.NewUnsignedInteger(len(slice))
}
