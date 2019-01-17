package main

import (
    "github.com/gorilla/mux"
)

// Funcao NewRouter modificada para evitar repeticao de codigo
func NewRouter() *mux.Router {
    router := mux.NewRouter()
    router.StrictSlash(true)

    // Percorre todas as routes criadas em routes.go, setando Methods e HandleFunc's
    for _, route := range routes {
        handler := route.HandlerFunc

        router.
        Methods(route.Method).
        Path(route.Pattern).
        Name(route.Name).
        Handler(handler)
    }
    return router
}
