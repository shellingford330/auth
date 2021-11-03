package main

import "log"

func main() {
	server := initializeServer()
	log.Fatalf("failed to linsten and serve. %+v", server.Start())
}
