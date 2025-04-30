package main

import (
	"fmt"

	"github.com/msisdev/dotato/dotato"
)

func main() {
	dtt := dotato.New()

	base, _, err := dtt.GetGroupBase("example", "nux")
	if err != nil {
		panic(err)
	}

	es, err := dtt.GetImportPaths("example", base)
	if err != nil {
		panic(err)
	}

	for _, e := range es {
		fmt.Printf("%s\n", e.Path.Abs())
	}
}
