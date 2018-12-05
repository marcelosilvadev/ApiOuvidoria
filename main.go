package main

import (
	"log"

	"bitbucket.org/ApiOuvidoria/api"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) //Informa o local do erro
	api := api.App{}
	api.StartServer()
}
