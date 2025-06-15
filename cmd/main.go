package main

import "timenote/router"

func main() {
	r := router.SetUpRouter()
	r.Run()
}
