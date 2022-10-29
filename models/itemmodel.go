package models

import (
	"database/sql"
	"fmt"

	"github.com/aliffathurrisqi/Golang-Iventory/config"
	"github.com/aliffathurrisqi/Golang-Iventory/entities"
)

type ItemModel struct{
	conn *sql.DB
}

func NewItemModel() *ItemModel{
	conn, err := config.DBconnection()

	if err != nil{
		panic(err)
	}

	return &ItemModel{
		conn: conn,
	}
}

func (p *ItemModel) All() ([]entities.Item, error){

	rows, err := p.conn.Query("SELECT * FROM items")

	if err != nil{
		return []entities.Item{}, err
	}

	defer rows.Close()

	var dataItem []entities.Item
	for rows.Next(){
		var item entities.Item
		rows.Scan(&item.Id, &item.Name, &item.Type_id, &item.Price, &item.Stock)

		dataItem = append(dataItem, item)
	}

	return dataItem, nil
}


func (p *ItemModel) Create(item entities.Item) bool{
	result, err := p.conn.Exec("INSERT INTO items VALUES(NULL,?,?,?,?)", 
	item.Name, item.Type_id, item.Price, item.Stock)

	if err != nil{
		fmt.Println(err)
		return false
	}

	lastInsertID, _ := result.LastInsertId()

	return lastInsertID > 0
}