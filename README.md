# take5

library in go by takemon-go

## Installation

```
$ go get -u github.com/takemon-go/take5
```

## Examples

### ReadLines

```
$ cat dog/main.go
package main

import (
	"fmt"
	"github.com/takemon-go/take5"
)

func main() {
	take5.ReadLines(func(line string) {
		fmt.Println(line)
	})
}
```

```
$ cat dog/main.go | go run ./dog
```

```
$ go run ./dog dog/main.go
```
