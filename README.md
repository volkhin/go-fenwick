# Fenwick Tree
[![GoDoc](https://godoc.org/github.com/volkhin/go-fenwick?status.svg)](https://godoc.org/github.com/volkhin/go-fenwick)
[![Build Status](https://travis-ci.org/volkhin/go-fenwick.svg?branch=master)](https://travis-ci.org/volkhin/go-fenwick)

Simple but still useful data structure [Fenwick
Tree](https://en.wikipedia.org/wiki/Fenwick_tree) implemented in Go language.

Install package with
```bash
go get github.com/volkhin/go-fenwick
```

Import it with
```go
import "github.com/volkhin/go-fenwick"
```

Use `fenwick` package name inside the code.

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
