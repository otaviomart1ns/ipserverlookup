package main

import (
	"fmt"
	"ipserverlookup/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Result:")

	aplication := app.Create()
	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
