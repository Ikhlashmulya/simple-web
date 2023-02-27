package handler 

import(
    "net/http"
    "html/template"
)

func Home(writer http.ResponseWriter, request *http.Request) {
    
    if request.URL.Path != "/" {
        http.NotFound(writer, request)
        return
    }
    
    temp, err := template.ParseFiles("views/index.html", "views/layout.html")
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
    
    temp.Execute(writer, nil)
}

func About(writer http.ResponseWriter, request *http.Request) {
    temp, err := template.ParseFiles("views/about.html", "views/layout.html")
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
    
    temp.Execute(writer, nil)
}

func Contact(writer http.ResponseWriter, request *http.Request) {
    temp, err := template.ParseFiles("views/contact.html", "views/layout.html")
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
    
    temp.Execute(writer, nil)
}