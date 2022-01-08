package main

import (
	"fmt"
)

type Settings struct {
	host string
	port int
}

type Server struct {
	Settings
}

func createServer(options ...func(s *Settings)) Server {
	settings := &Settings{
		host: "localhost",
		port: 8080,
	}

	for _, option := range options {
		option(settings)
	}

	return Server{
		Settings: *settings,
	}
}

func main() {
	configureHost := func(s *Settings) {
		s.host = "somehost"
	}

	configurePort := func(s *Settings) {
		s.port = 2121
	}

	server := createServer(configureHost, configurePort)

	fmt.Print(server)
}
