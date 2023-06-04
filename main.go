package main

import (
	"btpn-finpro/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8000")
}
