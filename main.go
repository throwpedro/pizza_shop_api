package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/json/products", getProducts)
	router.POST("/json/products", createProducts)

	router.Run("localhost:8080")
}

type product struct {
	ID        string  `json:"id"`
	Name      string  `json:"title"`
	Count     int     `json:"count"`
	UnitPrice float64 `json:"unitPrice"`
	Topping   string  `json:"topping"`
}

var products = []product{
	{ID: "1", Name: "Pepperoni", Count: 0, UnitPrice: 10, Topping: "Pepperoni"},
	{ID: "2", Name: "Cheese", Count: 0, UnitPrice: 10, Topping: "Cheese"},
	{ID: "3", Name: "Sausage", Count: 0, UnitPrice: 12, Topping: "Sausage"},
	{ID: "4", Name: "Veggie", Count: 0, UnitPrice: 15, Topping: "Veggie"},
	{ID: "5", Name: "Meat Lovers", Count: 0, UnitPrice: 20, Topping: "Meat Lovers"},
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func createProducts(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}
