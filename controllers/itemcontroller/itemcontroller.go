package itemcontroller

import (
	"html/template"
	"net/http"

	"github.com/aliffathurrisqi/Golang-Iventory/entities"
	"github.com/aliffathurrisqi/Golang-Iventory/models"
)

var itemModel = models.NewItemModel()

func Index(response http.ResponseWriter, request *http.Request){

	
	if request.Method == http.MethodGet{

	temp, err := template.ParseFiles("views/item/index.html")

	if err != nil{
		panic(err)
	}

	temp.Execute(response, nil)

	} else if request.Method == http.MethodPost{
		request.ParseForm()

		var item entities.Item

		item.Name = request.Form.Get("name")
		item.Type_id = request.Form.Get("type_id")
		item.Price= request.Form.Get("price")
		item.Stock = request.Form.Get("stock")

		itemModel.Create(item)

		http.Redirect(response, request, "/item", http.StatusFound)
	}

}

func Add(response http.ResponseWriter, request *http.Request){
	
}

func Edit(response http.ResponseWriter, request *http.Request){

}

func Delete(response http.ResponseWriter, request *http.Request){
	
}