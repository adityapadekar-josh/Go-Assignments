# Race Condition in Go: Example and Solution

## Problem

In the below code snippet concurrent goroutines execution corrupts a piece of data by
accessing it simultaneously it leads in raise condition.

You can view and run the code here: 
```go
package main

import (
	"fmt"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0

	go func() {
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() {
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
```

### Example Output

When you run the given code, the output may appear as:

```
1 is Even // (may vary when you run multiple times)
```

This output is incorrect and unpredictable, as the concurrent goroutines are corrupting the data during execution.

### Expected Output

The correct output should always be:

```
0 is Even
```

Update the given code to print the correct output.