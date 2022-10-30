package models

import (
	"database/sql"

	"github.com/aliffathurrisqi/Golang-Iventory/config"
	"github.com/aliffathurrisqi/Golang-Iventory/entities"
)

type TypeModel struct{
	conn *sql.DB
}

func NewTypeModel() *TypeModel{
	conn, err := config.DBconnection()

	if err != nil{
		panic(err)
	}

	return &TypeModel{
		conn: conn,
	}
}

func (p *TypeModel) All() ([]entities.Type, error){

	rows, err := p.conn.Query("SELECT * FROM types")

	if err != nil{
		return []entities.Type{}, err
	}

	defer rows.Close()

	var dataType []entities.Type
	for rows.Next(){
		var types entities.Type
		rows.Scan(&types.Id, &types.Name)

		dataType = append(dataType, types)
	}

	return dataType, nil
}