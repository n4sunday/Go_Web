package main
import (
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
)
type Product struct{
    Name string
    Price int
}
func main() {
    var templates = template.Must(template.ParseFiles("index.html"))
    router:=mux.NewRouter()
    router.HandleFunc("/",func(res http.ResponseWriter, req *http.Request){
        myProduct := Product{"นมสด",250}
        templates.ExecuteTemplate(res,"index.html",myProduct)
    })

    http.ListenAndServe(":8080",router)
}