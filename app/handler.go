package main

import (
	"errors"
	"io"
	"log/slog"
	"net"

	"github.com/codecrafters-io/redis-starter-go/parser"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
		}

		command := parser.Decode(buffer[:n])
		slog.Info("received command", "command", command)

		response := command.Run()

		slog.Info("responding to client", "response", response)

		conn.Write([]byte(response))
	}
}
