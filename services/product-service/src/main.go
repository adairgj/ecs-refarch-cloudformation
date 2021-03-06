package main

import "github.com/gin-gonic/gin"

type product struct {
	ID          string
	Title       string
	Description string
	Price       float64
}

func main() {

	router := gin.Default()

	// Respond to GET requests
	router.GET("/products", func(c *gin.Context) {

		products := []product{
			product{
				ID:          "0000-0000-0001",
				Title:       "Cialis",
				Description: "Mens Health - used to treat ED and BPH",
				Price:       300,
			},
			product{
				ID:          "0000-0000-0002",
				Title:       "Humalog",
				Description: "Diabetes - fast acting insulin used to control blood surgar in type 1 or type 2 diabetes ",
				Price:       600,
			},
			product{
				ID:          "0000-0000-0003",
				Title:       "Taltz",
				Description: "Immunology - used to treat adults with moderate to severe plaque psoriasis",
				Price:       14000,
			},
		}

		c.IndentedJSON(200, products)

	})

	// Serve all of the things..
	router.Run(":8001")

}
