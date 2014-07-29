# Fenwick Tree
Simple but still useful data structure [Fenwick
Tree](https://en.wikipedia.org/wiki/Fenwick_tree) implemented in Go language.

`go get` it as `github.com/volkhin/go-fenwick`, import it as
`github.com/volkhin/go-fenwick`, use it as `fenwick`.

## Usage example

```go
package main

import (
	"fmt"

	"github.com/volkhin/go-fenwick"
)

func main() {
	f := fenwick.New(10)
	f.Add(1, 4)
	f.Add(3, 123)
	fmt.Println(f.Sum(1, 10)) // 127
}
```

## License

MIT (c) 2014 Artem Volkhin
