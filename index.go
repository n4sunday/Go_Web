package main
import (
    "net/http"
    "fmt"
)
func main() {
    http.HandleFunc("/",index)
    http.HandleFunc("/login",login)
    http.ListenAndServe(":8080",nil)
}
func index(res http.ResponseWriter, req *http.Request){
    http.ServeFile(res,req,"index.html")
}
func login(res http.ResponseWriter, req *http.Request){
    fmt.Println("Method:",req.Method)
    req.ParseForm()
    fmt.Println("UserName:",req.Form["username"])
    fmt.Println("Password:",req.Form["password"])

}