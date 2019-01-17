package main

import (
    "net/http"
    "fmt"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
type Routes []Route

// Aqui criamos as routes possiveis e definimos suas HandlerFunc
var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "GetProducts",
        "GET",
        "/products",
        GetProducts,
    },
    Route{
        "GetProduct",
        "GET",
        "/products/{id}",
        GetProduct,
    },
}

// Aqui criamos as HandlerFunc para cada route
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Va para a URL '/products' para ver a lista de produtos ou '/products/{id}' para ver um produto especifico.")
    fmt.Fprintln(w, "IDs vao de 000 ate 025.")
}

// Mostra todos os produtos existentes
func GetProducts(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(products)
}

// Mostra um unico produto
func GetProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    for _, item := range products {
        if item.ID == vars["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
}
