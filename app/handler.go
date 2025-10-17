package main

import (
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
			slog.Error("reading bytes from connection", "err", err.Error())
		}

		if n == 0 {
			conn.Write([]byte("no data received"))
		}

		command := parser.Decode(buffer)
		response := command.Run()

		conn.Write([]byte(response))
	}
}
