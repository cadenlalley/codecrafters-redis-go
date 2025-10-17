package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
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

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("accepting connection", "err", err.Error())
		}

		go handleConnection(conn)
	}
}
