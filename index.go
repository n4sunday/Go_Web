package main
import (
    "fmt"
    "net/http"
)
func main() {
    http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request){
        fmt.Fprintf(res,"Nattapon")
    })
    http.HandleFunc("/product",product)

    http.ListenAndServe(":8080",nil)
}

func product(res http.ResponseWriter, req *http.Request){
        fmt.Fprintf(res,"Product")
}