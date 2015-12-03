# exit

A elegant way to exit.



## Usage

``` go
package main

import (
	"fmt"
	"time"

	"github.com/gocubes/exit"
)

func main() {
	fmt.Println("Start some service.")
	// Wait interrupt signal, and set things number(0-255) should todo before exit.
	exit.Wait(0)

	for {
		time.Sleep(1e6)
		// Received an interrupt signal
		if exit.Exit == true {
			fmt.Println("Clean some things")
			// Done.
			exit.Todo.Done()
			break
		}
	}

}
```

