package utils

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/kamalesh-servi/web-scraper/models"
)

func CsvData() {
	// opening the CSV file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// defining the CSV headers
	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	// writing the column headers
	writer.Write(headers)

	// adding each Pokemon product to the CSV output file
	for _, pokemonProduct := range models.Products {
		// converting a PokemonProduct to an array of strings
		record := []string{
			pokemonProduct.Url,
			pokemonProduct.Image,
			pokemonProduct.Name,
			pokemonProduct.Price,
		}

		// writing a new CSV record
		err := writer.Write(record)
		if err != nil {
			panic(err)
		}
	}
	defer writer.Flush()

}
