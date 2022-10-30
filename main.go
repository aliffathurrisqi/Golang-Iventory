package main

import (
	"net/http"

	itemcontroller "github.com/aliffathurrisqi/Golang-Iventory/controllers"
)

func main(){

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/item", itemcontroller.Index)

	http.HandleFunc("/item/add", itemcontroller.Add)

	http.HandleFunc("/item/edit", itemcontroller.Edit)

	http.HandleFunc("/item/delete", itemcontroller.Delete)

	http.ListenAndServe(":3000", nil)
}