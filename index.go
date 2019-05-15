package main
import (
    "fmt"
    "net/http"
)
func main() {
    userDB:=map[string]int{
        "Sun":21,
        "Ding":23,
        "Noey":25,
    }

    http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request){
        fmt.Fprintf(res,"Server Go Ready")
    })

    http.HandleFunc("/user/",func(res http.ResponseWriter, req *http.Request){
        name:=req.URL.Path[len("/user/"):]
        age:=userDB[name]
        fmt.Fprintf(res,"My Name is %s %d year old",name,age)
    })



    http.ListenAndServe(":8080",nil)
}