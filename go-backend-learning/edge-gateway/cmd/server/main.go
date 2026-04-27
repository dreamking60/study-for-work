package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("edge-gateway starting on :8080")
	
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:	":8080",
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

}

