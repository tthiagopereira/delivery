package main

import (
	"fmt"
	route2 "github.com/tthiagopereira/delivery/application/route"
)

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPosition()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}
