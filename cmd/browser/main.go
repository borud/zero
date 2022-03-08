package main

import (
	"context"
	"log"
	"time"

	"github.com/grandcat/zeroconf"
)

func main() {
	for {
		triggerBrowse()
		time.Sleep(1 * time.Second)
	}
}

func triggerBrowse() {
	entries := make(chan *zeroconf.ServiceEntry)

	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			log.Printf("> %+v\n", entry)
		}
		log.Println("No more entries.")
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	err = resolver.Browse(ctx, "_myservice._tcp", "local.", entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err)
	}
	<-ctx.Done()
}
