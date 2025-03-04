package main

import (
	"fmt"
	"log"
	"net/http"

	sv_cfg "backend/servercfg"
)

func main() {
	sv_cfg.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola mundo!")
	})

	fmt.Print("Iniciando servidor web\n")
	log.Fatal(sv_cfg.Server.ListenAndServe())

}
