package main
import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)
func main() {
    userDB:=map[string]int{
        "Sun":21,
        "Ding":23,
        "Noey":25,
    }
    router:=mux.NewRouter()
    router.HandleFunc("/",func(res http.ResponseWriter, req *http.Request){
        fmt.Fprintf(res,"Server Go Ready")
    })

    router.HandleFunc("/user/{name}",func(res http.ResponseWriter, req *http.Request){
        vars:=mux.Vars(req)
        name:=vars["name"]
        age:=userDB[name]
        fmt.Fprintf(res,"My Name is %s %d year old",name,age)
    }).Methods("GET")
    
    http.ListenAndServe(":8080",router)
}