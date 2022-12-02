package models

import (
	"github.com/furtadomn/golang-studies/Go_Crie_uma_Aplicacao_Web/db"
)

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func FindAllProducts() []Product {
	db := db.DBConnect()

	selectAllProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		products = append(products, p)
	}

	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int64) {
	db := db.DBConnect()

	insertData, err := db.Prepare("insert into products(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)
	defer db.Close()
}
