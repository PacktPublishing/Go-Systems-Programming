package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("Got", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			sig := <-sigs
			fmt.Println(sig)
			handleSignal(sig)
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(10 * time.Second)
	}
}
