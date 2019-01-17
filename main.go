package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    // Utiliza a funcao NewRouter() que eh criada em router.go
    router := NewRouter()

    // Nao e necessario definir a PORT, ja que o Heroku escolhe uma dinamicamente
    // A port fica armazenada em uma variavel de ambiente, que podemos ler
    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}
