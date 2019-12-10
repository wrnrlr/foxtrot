package main

import (
	"github.com/wrnrlr/foxtrot/app"
	"os"
)

func main() {
	path := ""
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	app.RunUI(path)
}
