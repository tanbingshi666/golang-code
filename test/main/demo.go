package main

import (
	"fmt"
	"golang-code/test/utils"
)

var xi = utils.Hello

func init() {
	fmt.Printf("init %s \n", xi)
}

func main() {
	fmt.Println("main")

}
