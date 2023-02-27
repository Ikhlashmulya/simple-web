package main

import (
    "log"
    "net/http"
    "github.com/Ikhlashmulya/simple-web/handler"
)

func main() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", handler.Home)
    mux.HandleFunc("/about", handler.About)
    mux.HandleFunc("/contact", handler.Contact)
    
    fileServer := http.FileServer(http.Dir("assets"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))
    
    log.Println("Starting web on port 8080")
    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}