package main

import (
	"fmt"
	"log"
	"os"
	"search-web/app"
)

func main() {
	fmt.Println("# Search Web")

	aplicacao := app.Generate()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
