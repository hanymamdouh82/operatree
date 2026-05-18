package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hanymamdouh82/operatree/internal/project"
)

func main() {
	fmt.Println("Welcome OperaTree...A project operating system built on your filesystem...!")

	if len(os.Args) < 3 {
		log.Fatal(fmt.Errorf("not enough args"))
	}

	cmd := os.Args[1]

	switch cmd {
	case "bootstrap":
		v := os.Args[2]
		fmt.Printf("i will bootstrap for project %s\n", v)
		project.Bootstrap("/mnt/extra/onfly", v)
	case "exit":
		fmt.Println("existing")
		os.Exit(0)
	default:
		fmt.Println("unknown command")
	}
}
