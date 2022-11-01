package main

import (
	"fmt"
	"davidone.it/singers/utils"
	"davidone.it/singers/api"
)

func main() {
    fmt.Println("Singer Application Started!")
	fmt.Println(utils.ToUpper("Hello, World!"))
	api.Router()
	fmt.Println("Singer Application Stopping!")
}