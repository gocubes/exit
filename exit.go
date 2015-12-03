package exit

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

type debug bool

var (
	Debug      debug
	Exit       bool
	Todo       sync.WaitGroup
	signalChan = make(chan os.Signal, 1)
)

// Print debug info.
func (d debug) println(args ...interface{}) {
	if d {
		fmt.Println(args...)
	}
}

// Waiting for exist.
func Wait(delta uint8) {
	Debug.println("Will waiting for exist.")
	Todo.Add(int(delta))
	go func() {
		signal.Notify(signalChan, os.Interrupt)
		for range signalChan {
			Debug.println("\nReceived an interrupt, stopping services...")
			Exit = true
			Todo.Wait()
			os.Exit(0)
		}
	}()

}
