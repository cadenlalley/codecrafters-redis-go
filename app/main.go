package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"strings"
)

const ADDRESS = "0.0.0.0:6379"

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	err := run()
	if err != nil {
		slog.Error("starting server", "err", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	slog.Info("starting server", "address", ADDRESS)

	listener, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		return fmt.Errorf("binding to address: %s", ADDRESS)
	}

	conn, err := listener.Accept()
	if err != nil {
		slog.Error("accepting connection", "err", err.Error())
	}

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			panic(err)
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
