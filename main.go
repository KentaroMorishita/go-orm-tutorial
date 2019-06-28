package main

import (
	"fmt"
	"log"
	"net/http"

	"go-orm-tutorial/routes"
)

func main() {
	fmt.Println("Go ORM Tutorial")
	router := routes.Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}
