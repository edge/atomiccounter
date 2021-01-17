# Atomic Counter

This package provides a simple, encapsulated atomic counter with some basic calculator functions.

## Usage

This example demonstrates basic usage.

```go
package main

import (
	"fmt"

	counter "github.com/edge/atomiccounter"
)

func main() {
	// c *counter.Counter
	c := counter.New()

	// prints "50"
	c.Set(50)
	fmt.Println(c.Get())

	// prints "100"
	c.Add(50)
	fmt.Println(c.Get())

	// prints "101"
	c.Inc()
	fmt.Println(c.Get())
}
```

Refer to [atomiccounter.go](./atomiccounter.go) for all available functions on the Counter type.
