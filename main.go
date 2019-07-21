package main

import (
	"fmt"
	"time"

	"./router"
)

func main() {
	fmt.Println("h")
	currentTime := time.Now()
	fmt.Println("Year   :", currentTime.Year())
	fmt.Println("Month  :", currentTime.Month())
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080

}
