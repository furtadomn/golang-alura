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

	selectAllProducts, err := db.Query("select * from products order by id asc")
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

		p.Id = id
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

func DeleteProduct(id string) {
	db := db.DBConnect()

	deleteData, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteData.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.DBConnect()

	dbProduct, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for dbProduct.Next() {
		var (
			id, quantity      int
			name, description string
			price             float64
		)

		err = dbProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Nome = name
		productToUpdate.Descricao = description
		productToUpdate.Preco = price
		productToUpdate.Quantidade = quantity
	}

	defer db.Close()

	return productToUpdate
}

func UpdateProduct(id, quantity int64, name, description string, price float64) {
	db := db.DBConnect()

	updateProduct, err := db.Prepare("update products set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}
