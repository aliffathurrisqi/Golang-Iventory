package main

import (
	"net/http"

	_ "github.com/aliffathurrisqi/Golang-Iventory/controllers/itemcontroler"
)

func main(){
	http.HandleFunc("/", itemcontroller.index)
}