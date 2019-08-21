package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Cmd struct {
	Closed chan struct{}
}

func (c *Cmd) Close() {
	log.Println("closing program...")
	// wait for program to clean up nicely
	time.Sleep(5 * time.Second)
	log.Println("closed program")
	close(c.Closed)
}

// section: main
func main() {
	cmd := Cmd{Closed: make(chan struct{})}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	log.Println("Listening for signals")

	// Block until one of the signals above is received
	<-signalCh
	log.Println("Signal received, initializing clean shutdown...")
	go cmd.Close()

	// Block again until another signal is received, a shutdown timeout elapses,
	// or the Command is gracefully closed
	log.Println("Waiting for clean shutdown...")
	select {
	case <-signalCh:
		log.Println("second signal received, initializing hard shutdown")
	case <-time.After(time.Second * 30):
		log.Println("time limit reached, initializing hard shutdown")
	case <-cmd.Closed:
		log.Println("server shutdown completed")
	}

	// goodbye.
}

// section: main
