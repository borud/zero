package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/grandcat/zeroconf"
)

func main() {
	server, err := zeroconf.Register("hbgw", "_myservice._tcp", "local.", 7777, []string{"version=1.2.3", "room=foo", "floor=bar"}, nil)
	if err != nil {
		log.Fatalf("error registering: %v", err)
	}
	defer server.Shutdown()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}
