package main

import (
	"fmt"
	"log"

	"go-orm-tutorial/routes"
)

func main() {
	fmt.Println("Go ORM Tutorial")
	e := routes.Router()
	log.Fatal(e.Start(":8000"))
}
