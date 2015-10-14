
# GopherLibTerminal - Go Bindings for BearLibTerminal

This repository is a [Go](https://golang.org/) binding for the excellent [BearLibTerminal](http://foo.wyrd.name/en:bearlibterminal) library.

# Installing

## Requirements

[BearLibTerminal](http://foo.wyrd.name/en:bearlibterminal) is required to be on your system.

## Go Get

The usual `go get github.com/mpatraw/gopherlibterminal` should work.

# Usage

The official [reference](http://foo.wyrd.name/en:bearlibterminal:reference) for [BearLibTerminal](http://foo.wyrd.name/en:bearlibterminal) should serve as the main documentation. All functions and constants have been renamed to honor [Go](https://golang.org/)'s naming convention. `terminal_open` becomes `gopherlibterminal.Open`.

The package name is shortened to "glt" for convenience.

```go
package main

import "github.com/mpatraw/gopherlibterminal"

func main() {
	glt.Open()
	defer glt.Close()
	glt.Print(0, 0, "Hello, world!")
	glt.Refresh()
	glt.Delay(1000)
}
```

# Documentation

See [BearLibTerminal](http://foo.wyrd.name/en:bearlibterminal)'s official [reference](http://foo.wyrd.name/en:bearlibterminal:reference).

# Distribution

In order for your executable to be run, you will have to package the appropriate [BearLibTerminal](http://foo.wyrd.name/en:bearlibterminal) shared library for a specific architecture.
