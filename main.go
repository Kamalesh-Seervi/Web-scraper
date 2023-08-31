package main

import (

	// "log"

	"github.com/gocolly/colly"
	"github.com/kamalesh-servi/web-scraper/models.go"
)

func main() {
	// defining a data structure to store the scraped data
	// initializing the slice of structs that will contain the scraped data
	var pokemonProducts []models.PokemonProduct
	// initialize a slice of PokemonProduct that will contain the scraped data:
	c := colly.NewCollector()
	c.Visit("https://scrapeme.live/shop/")

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		// initializing a new PokemonProduct instance
		pokemonProduct := models.PokemonProduct{}

		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)

	})

}
