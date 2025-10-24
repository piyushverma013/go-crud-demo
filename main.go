package main

import "github.com/piyushverma013/simple-go-crud/router"

func main() {
	r := router.SetUpRouter()
	r.Run(":8080")
}
