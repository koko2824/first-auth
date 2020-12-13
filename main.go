package main

import (
	"first-auth/db"
	"first-auth/router"
)

func main() {
	db.Init()
	router.Init()
}
