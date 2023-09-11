package main

import (
	"log"

	"github.com/nikitamirzani323/btangkas-client/router"
)

func main() {
	app := router.Init()
	log.Fatal(app.Listen(":2116"))
}
