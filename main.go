package main

import (
	"fmt"

	"./router"
)

func main() {
	fmt.Println("h")
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080

}
