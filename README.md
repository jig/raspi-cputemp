# raspi-cputemp

Minimal package to read ARM CPU of the Raspberry Pi 4.

```go
package main

import (
	"log"

	cputemp "github.com/jig/raspi-cputemp"
)

func main() {
	temp, err := cputemp.Get()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CPU temp is %f°C", temp)
}
```