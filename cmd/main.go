package main

import (
	"timenote/db"
	"timenote/router"
)

func main() {
	db.InitDB()
	r := router.SetUpRouter()
	r.Run()
}
