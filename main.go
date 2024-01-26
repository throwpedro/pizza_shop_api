package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.GET("/json/products", getProducts)
	router.POST("/json/products", createProducts)
	router.PATCH("/json/products", updateProducts)
	router.DELETE("/json/products/:id", deleteProduct)

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
	c.JSON(http.StatusOK, products)
}

func createProducts(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProducts(c *gin.Context) {
	var updatedProduct product

	if err := c.BindJSON(&updatedProduct); err != nil {
		return
	}

	found := false
	for i, p := range products {
		if p.ID == updatedProduct.ID {
			products[i] = updatedProduct
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			break
		}
	}
}
