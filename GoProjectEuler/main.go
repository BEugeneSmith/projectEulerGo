package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"GoProjectEuler/euler"
)

func main() {

	name := os.Args[1]
	arg, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	RoutesFive(name, arg)

}

// RoutesFive returns a function
func RoutesFive(s string, i int) {
	switch s {
	case "euler1":
		fmt.Println(euler.Euler1(i))
	}
}
