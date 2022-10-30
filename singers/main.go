package main

import (
	"fmt"
	"davidone.it/singers/utils"
	"davidone.it/singers/api"
)

func main() {
    fmt.Println("Hello, World!")
	fmt.Println(utils.ToUpper("Hello, World!"))
	api.Router()
}