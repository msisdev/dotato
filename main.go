package main

import "github.com/msisdev/dotato/dotato"

func main() {
	dtt := dotato.New()

	base, _, err := dtt.GetGroupBase("example", "nux")
	if err != nil {
		panic(err)
	}

	err = dtt.GetImportPaths("example", base)
	if err != nil {
		panic(err)
	}
}
