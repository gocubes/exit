package main

import (
	"fmt"
	"time"

	"github.com/gocubes/exit"
)

func main() {
	fmt.Println("Start some service.")
	exit.Wait(0)

	for {
		time.Sleep(1e6)
		if exit.Exit == true {
			fmt.Println("Clean some things")
			exit.Todo.Done()
			break
		}
	}

}
