package main

import (
	"fmt"
	"net/http"

	"github.com/rahulp18/url-shortner/handler"
	"github.com/rahulp18/url-shortner/service"
	"github.com/rahulp18/url-shortner/store"
)

func main() {
	memStore := store.NewMemoryStore()
	urlService := service.NewURLService(memStore)
	urlHandler := handler.NewURLHandler(urlService)

	http.HandleFunc("/shorten", urlHandler.Shorten)
	http.HandleFunc("/", urlHandler.Redirect)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)

}
