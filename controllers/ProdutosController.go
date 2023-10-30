package controllers

import (
	"database/sql"
	"golangMysql/config"
	entities "golangMysql/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	db *sql.DB = config.ConfigDb()
)

func ListProduto(ctx *gin.Context) {
	slicP := []entities.Produto{}

	res, err := db.Query("SELECT * FROM produtos")
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var produtos entities.Produto

		err := res.Scan(&produtos.ID, &produtos.Nome, &produtos.Preco, &produtos.Quantidade)

		if err != nil {
			log.Fatal(err)
		}
		slicP = append(slicP, produtos)
	}

	ctx.JSON(http.StatusOK, slicP)
}

func CreateProduto(ctx *gin.Context) {
	nProd := entities.NewProduto()
	if err := ctx.BindJSON(nProd); err != nil {
		log.Fatal(err)
	}
	res, err := db.Prepare("INSERT INTO produtos(idprodutos,nome,preco,quantidade) VALUES (?,?,?,?)")
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}
	_, err = res.Exec(nProd.ID, nProd.Nome, nProd.Preco, nProd.Quantidade)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "Produto criado com sucesso!",
	})
}
func seachProduto(id *string) bool {

	res, err := db.Prepare("SELECT * FROM produtos WHERE idprodutos = ?")

	if err != nil {
		log.Fatal(err)
	}

	item, err := res.Exec(*id)

	if err != nil {
		log.Fatal(err)
	}
	if item != nil {
		return true
	}
	return false

}

func UpdateProduto(ctx *gin.Context) {

	Idparams := ctx.Param("id")

	var body entities.Produto
	if seachProduto(&Idparams) != false {
		res, err := db.Prepare("UPDATE produtos SET nome = ?,preco = ?,quantidade = ? WHERE idprodutos = ?")
		if err != nil {
			panic(err)
		}

		if err := ctx.BindJSON(&body); err != nil {
			panic(err)
		}
		_, err = res.Exec(body.Nome, body.Preco, body.Quantidade, Idparams)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusAccepted, gin.H{
			"Message": "Produto Alterado!",
		})
	}
}

func DeleteProduto(ctx *gin.Context) {
	Idparam := ctx.Param("id")
	if seachProduto(&Idparam) != false {
		res, err := db.Prepare("DELETE FROM produtos WHERE idprodutos = ?")

		if err != nil {
			panic(err)
		}

		_, err = res.Exec(Idparam)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Produto Deletado!",
		})
	}
}
