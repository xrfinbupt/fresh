package main

import (
	"github.com/c2h5oh/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world - Martini"
	})
	m.Run()
}
