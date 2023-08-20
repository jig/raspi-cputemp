# raspi-cputemp

Minimal package to read ARM CPU of the Raspberry Pi 4.

Idea from https://github.com/fgrosse/pi-temp.

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
	log.Printf("CPU temp is %.3f°C", temp)
}
```

# Command line

```bash
cd cmd/temp
go build
go install

temp
```

Example output:

```bash
2023/08/20 15:08:38 CPU temp is 84.237°C
```

A bit too high temperature by the way.
