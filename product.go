package main

import (
    "bufio"
    "io"
    "os"
    "log"
    "encoding/csv"
)

// Definicao de um product
type Product struct {
    ID    string `json:"id,omitempty"`
    Name  string `json:"name,omitempty"`
    Price string `json:"price,omitempty"`
}
var products []Product

// Popula products com dados de um arquivo csv
func init() {
    csvFile,_ := os.Open("product_db.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        products = append(products, Product{
            ID: line[0],
            Name: line[1],
            Price: line[2],
        })
    }
}
