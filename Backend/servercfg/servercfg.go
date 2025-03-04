package servercfg

import "net/http"

var (
	Mux = http.NewServeMux()

	Server = &http.Server{
		Addr:    "localhost:4451",
		Handler: Mux,
	}
)
