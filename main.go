package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/kamalesh-servi/web-scraper/models"
	"github.com/kamalesh-servi/web-scraper/utils"
)

func main() {
	// defining a data structure to store the scraped data
	// initializing the slice of structs that will contain the scraped data
	// initialize a slice of PokemonProduct that will contain the scraped data:
	c := colly.NewCollector()
	c.Visit("https://scrapeme.live/shop/")

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := models.PokemonProduct{}
		pokemonProduct.Url = e.ChildAttr("a", "href")
		pokemonProduct.Image = e.ChildAttr("img", "src")
		pokemonProduct.Name = e.ChildText("h2")
		pokemonProduct.Price = e.ChildText(".price")

		models.Products = append(models.Products, pokemonProduct)
	})
	c.Wait()

	utils.CsvData()

	fmt.Println(models.Products)

}
