## Detect connected android devices

Example:
```
package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/dimorinny/android-devices"
	"log"
)

func main() {
	devices, err := android.Devices()
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(devices)
}

```