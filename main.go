package main

import (
    "html/template"
    "log"
    "net/http"
)

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("./templates/home.html")
        if err != nil {
            log.Fatal("Error loading template: ", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        err = tmpl.Execute(w, nil)
        if err != nil {
            log.Fatal("Error executing template: ", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        }
    })

    log.Println("Starting server on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
