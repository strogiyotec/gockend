package main

import (
	"gockend/controllers"
)

func main() {
	router := controllers.CreateMapping()
	router.Run()
}
