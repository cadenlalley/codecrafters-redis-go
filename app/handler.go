package main

import (
	"log/slog"
	"net"
	"strings"
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

		bufferString := string(buffer)
		parts := strings.Split(bufferString, "\r\n")
		if parts[0] == "*" {
			parts = parts[1:]
		}

		for i := 2; i < len(parts); i += 2 {
			if parts[i] == "PING" {
				conn.Write([]byte("+PONG\r\n"))
			}
		}
	}
}
